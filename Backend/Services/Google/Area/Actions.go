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

// Google Actions
// @Summary send all the Actions
// @Description send all the Actions available on the Google services as an object arrays with the names and the object needed
// @Tags Google Area
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string "Response is the received data"
// @Failure 400 {object} map[string]string "Invalid request it contains the error"
// @Failure 500 {object} map[string]string "Internal error it contains the error"
// @Router /actions [get]
func GetActions(c *gin.Context) {
	b, err := utils.OpenFile("Models/Actions.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	json := utils.BytesToJson(b)
	c.JSON(http.StatusOK, json)
}

// Google Services
// @Summary Register an received Actions
// @Description Register the Actions received by the message brocker with all informations nedded
// @Tags Google Area
// @Accept json
// @Produce json
// @Param routes body models.ReceivedActions true "It must contains the AreaId and the reactions type"
// @Success 200 {object} map[string]string "Response is the received data"
// @Failure 400 {object} map[string]string "Invalid request it contains the error"
// @Failure 500 {object} map[string]string "Internal error it contains the error"
// @Router /action [post]
func StoreActions(c *gin.Context) {
	var receivedData models.ReceivedActions
	var user models.User
	db := utils.OpenDB(c)

	if err := c.ShouldBindJSON(&receivedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data", "details": err.Error()})
		return
	}

	id := utils.ParseToken(receivedData.UserToken)
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token"})
		return
	}

	row := db.QueryRow(c, "SELECT * FROM \"User\" WHERE id = $1", id)
	err := row.Scan(&user.Id, &user.Mail, &user.Password, &user.Login, &user.Lastname, &user.AsanaToken, &user.DiscordToken,
		&user.DropboxToken, &user.GithubToken, &user.GitlabToken, &user.GoogleToken, &user.MiroToken, &user.SpotifyToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	gmailProfile, err := GetGmailProfile(*user.GoogleToken)

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

	db.QueryRow(context.Background(), query, *user.GoogleToken, receivedData.AreaId, receivedData.ActionType, gmailProfile.MessagesTotal, receivedData.NbEvents).Scan(&lastInsertID)
	defer db.Close(c)

	c.JSON(http.StatusOK, gin.H{"message": "Action stored successfully", "id": lastInsertID})
}
