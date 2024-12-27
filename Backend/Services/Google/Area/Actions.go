package area

import (
	models "google/Models"
	"google/utils"
	"io"
	"net/http"

	"context"

	"github.com/gin-gonic/gin"
)

func GetNbEvents(token string) int {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://www.googleapis.com/calendar/v3/calendars/primary/events", nil)

	if err != nil {
		return -1
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Accept", "application/json")

	resp, _ := client.Do(req)
	b, _ := io.ReadAll(resp.Body)
	json := utils.BytesToJson(b)
	if json == nil || json["items"] == nil {
		return -1
	}
	defer resp.Body.Close()
	return len(json["items"].([]any))
}

func StoreActions(c *gin.Context) {
	var receivedData models.ReceivedActions
	db := utils.OpenDB(c)

	if err := c.ShouldBindJSON(&receivedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data", "details": err.Error()})
		return
	}

	query := `
		INSERT INTO "GoogleActions" (user_token, area_id, action_type, nb_events)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	var lastInsertID int
	receivedData.NbEvents = GetNbEvents(receivedData.UserToken)

	err := db.QueryRow(context.Background(), query, receivedData.UserToken, receivedData.AreaId, receivedData.ActionType, receivedData.NbEvents).Scan(&lastInsertID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert data", "details": err.Error()})
		return
	}

	defer db.Close(c)

	c.JSON(http.StatusOK, gin.H{"message": "Action stored successfully", "id": lastInsertID})
}
