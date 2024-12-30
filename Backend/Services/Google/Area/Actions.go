package area

import (
	"encoding/json"
	"fmt"
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

func GetGmailProfile(accessToken string) (*models.GmailProfile, error) {
	url := "https://gmail.googleapis.com/gmail/v1/users/me/profile"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get Gmail profile: %s", resp.Status)
	}

	var profile models.GmailProfile
	if err := json.NewDecoder(resp.Body).Decode(&profile); err != nil {
		return nil, err
	}

	return &profile, nil
}

func StoreActions(c *gin.Context) {
	var receivedData models.ReceivedActions
	db := utils.OpenDB(c)

	if err := c.ShouldBindJSON(&receivedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data", "details": err.Error()})
		return
	}
	gmailProfile, err := GetGmailProfile(receivedData.UserToken)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data", "details": err.Error()})
		return
	}

	query := `
		INSERT INTO "GoogleActions" (user_token, area_id, action_type, nb_message, nb_events)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`

	var lastInsertID int
	receivedData.NbEvents = GetNbEvents(receivedData.UserToken)

	db.QueryRow(context.Background(), query, receivedData.UserToken, receivedData.AreaId, receivedData.ActionType, gmailProfile.MessagesTotal, receivedData.NbEvents).Scan(&lastInsertID)
	defer db.Close(c)

	c.JSON(http.StatusOK, gin.H{"message": "Action stored successfully", "id": lastInsertID})
}
