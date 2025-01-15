package area

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	models "google/Models"
	"google/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Google Reactions
// @Summary send all the Reactions
// @Description send all the Reactions available on the Google services as an object arrays with the names and the object needed
// @Tags Google Area
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string "Response is the received data"
// @Failure 400 {object} map[string]string "Invalid request it contains the error"
// @Failure 500 {object} map[string]string "Internal error it contains the error"
// @Router /reactions [get]
func GetReactions(c *gin.Context) {
	b, err := utils.OpenFile(models.ReactionsModelPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	json := utils.BytesToJson(b)
	c.JSON(http.StatusOK, json)
}

// Google Services
// @Summary Register an received Reactions
// @Description Register the reactions received by the message brocker with all informations nedded
// @Tags Google Area
// @Accept json
// @Produce json
// @Param routes body models.GoogleReaction true "It must contains the AreaId and the reactions type"
// @Success 200 {object} map[string]string "Response is the received data"
// @Failure 400 {object} map[string]string "Invalid request it contains the error"
// @Failure 500 {object} map[string]string "Internal error it contains the error"
// @Router /reaction [post]
func StoreReactions(c *gin.Context) {
	db := utils.OpenDB(c)

	defer db.Close(c)

	var receivedData models.GoogleReaction
	if err := c.ShouldBindJSON(&receivedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data", "details": err.Error()})
		return
	}

	query := `
		INSERT INTO "GoogleReactions" 
		(user_token, area_id, reaction_type, summary, description, start_time, end_time, attendees, recipient, subject, message)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
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
		receivedData.Recipient,
		receivedData.Subject,
		receivedData.Message,
	)

	c.JSON(http.StatusOK, gin.H{"message": "Reaction stored successfully"})
}

func CreateEvents(information models.GoogleReaction) (*http.Response, error) {
	body := &models.GoogleReactionSend{information.Summary, information.Description, models.DateTime{information.StartTime}, models.DateTime{information.EndTime}, []models.Attendee{{information.Attendees}}}
	client := &http.Client{}
	b, err := json.Marshal(body)

	if err != nil {
		return nil, err
	}

	r := bytes.NewReader(b)
	req, err := http.NewRequest("POST", "https://www.googleapis.com/calendar/v3/calendars/primary/events", r)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+information.UserToken)
	req.Header.Set("Content-Type", "application/json")

	return client.Do(req)
}

func SendEmail(reactions models.GoogleReaction) (*http.Response, error) {
	rawEmail := fmt.Sprintf(
		"To: %s\r\nSubject: %s\r\nContent-Type: text/plain; charset=\"UTF-8\"\r\n\r\n%s",
		reactions.Recipient, reactions.Subject, reactions.Message,
	)

	encodedEmail := base64.URLEncoding.EncodeToString([]byte(rawEmail))
	requestBody, err := json.Marshal(models.SendMessageRequest{Raw: encodedEmail})
	if err != nil {
		return nil, fmt.Errorf("failed to create request body: %v", err)
	}

	req, err := http.NewRequest("POST", "https://gmail.googleapis.com/gmail/v1/users/me/messages/send", strings.NewReader(string(requestBody)))
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+reactions.UserToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	return client.Do(req)
}

func FindReactions(id int, information models.GoogleReaction) (*http.Response, error) {
	reactions := map[int]func(models.GoogleReaction) (*http.Response, error){
		0: CreateEvents,
		1: SendEmail,
	}

	return reactions[id](information)
}

// Google Services
// @Summary Trigger an Area
// @Description Actions triggerd the reactions and call the trigger route
// @Tags Google trigger
// @Accept json
// @Produce json
// @Param routes body models.MessageBrocker true "It contains the Area Id to the reactions"
// @Success 200 {object} map[string]string "Response of the reactions"
// @Failure 400 {object} map[string]string "Invalid request it contains the error"
// @Failure 500 {object} map[string]string "Internal error it contains the error"
// @Router /trigger [post]
func Trigger(c *gin.Context) {
	receivedData := models.MessageBrocker{}
	information := models.GoogleReaction{}

	if err := c.ShouldBindJSON(&receivedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := utils.OpenDB(c)
	row := db.QueryRow(c, "SELECT user_token, reaction_type, summary, description, start_time, end_time, attendees, recipient, subject, message FROM \"GoogleReactions\" WHERE area_id = $1", receivedData.AreaId)

	if err := row.Scan(
		&information.UserToken,
		&information.ReactionType,
		&information.Summary,
		&information.Description,
		&information.StartTime,
		&information.EndTime,
		&information.Attendees,
		&information.Recipient,
		&information.Subject,
		&information.Message,
	); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer db.Close(c)

	rep, err := FindReactions(information.ReactionType, information)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(rep.StatusCode, gin.H{
		"body": rep.Body,
	})
}
