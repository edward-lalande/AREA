package area

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	models "google/Models"
	"google/utils"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func StoreReactions(c *gin.Context) {
	db := utils.OpenDB(c)

	defer db.Close(c)

	var receivedData models.GoogleCalendarReaction
	if err := c.ShouldBindJSON(&receivedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data", "details": err.Error()})
		return
	}
	fmt.Println("stored: ", receivedData)
	query := `
		INSERT INTO "GoogleReactions" 
		(user_token, area_id, reaction_type, summary, description, start_time, end_time, attendees)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`

	db.QueryRow(context.Background(), query,
		receivedData.UserToken,
		receivedData.AreaId,
		receivedData.ReactionType,
		receivedData.Summary,
		receivedData.Description,
		receivedData.StartTime,
		receivedData.EndTime,
		receivedData.Attendees,
	)

	c.JSON(http.StatusOK, gin.H{"message": "Reaction stored successfully"})
}

func CreateEvents(information models.GoogleCalendarReaction) (*http.Response, error) {
	body := &models.GoogleReactionSend{information.Summary, information.Description, models.DateTime{information.StartTime}, models.DateTime{information.EndTime}, []models.Attendee{{information.Attendees}}}
	client := &http.Client{}
	b, err := json.Marshal(body)

	if err != nil {
		return nil, err
	}

	json := utils.BytesToJson(b)
	fmt.Println("req: ", json)

	r := bytes.NewReader(b)
	req, err := http.NewRequest("POST", "https://www.googleapis.com/calendar/v3/calendars/primary/events", r)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+information.UserToken)
	req.Header.Set("Content-Type", "application/json")

	return client.Do(req)
}

func FindReactions(id int, information models.GoogleCalendarReaction) (*http.Response, error) {
	reactions := map[int]func(models.GoogleCalendarReaction) (*http.Response, error){
		0: CreateEvents,
	}

	return reactions[id](information)
}

func Trigger(c *gin.Context) {
	receivedData := models.MessageBrocker{}
	information := models.GoogleCalendarReaction{}

	if err := c.ShouldBindJSON(&receivedData); err != nil {
		fmt.Println("error: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("received: ", receivedData.AreaId)
	db := utils.OpenDB(c)
	row := db.QueryRow(c, "SELECT user_token, reaction_type, summary, description, start_time, end_time, attendees FROM \"GoogleReactions\" WHERE area_id = $1", receivedData.AreaId)

	if err := row.Scan(&information.UserToken, &information.ReactionType, &information.Summary, &information.Description, &information.StartTime, &information.EndTime, &information.Attendees); err != nil {
		fmt.Println("error on db : ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer db.Close(c)
	fmt.Println("reaction_type: ", information.ReactionType)
	fmt.Println("information: ", information)
	rep, err := FindReactions(information.ReactionType, information)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(rep.StatusCode, gin.H{
		"body": rep.Body,
	})
	b, _ := io.ReadAll(rep.Body)
	json := utils.BytesToJson(b)
	fmt.Println("rep: ", json)
}
