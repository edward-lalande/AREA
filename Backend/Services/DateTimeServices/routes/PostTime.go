package routes

import (
	"bytes"
	models "date-time-service/Models"
	"date-time-service/utils"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostTime(c *gin.Context) {
	var receivedData models.DataReceive
	if err := c.ShouldBindJSON(&receivedData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	db := utils.OpenDB(c)
	_, err := db.Exec(c, "INSERT INTO \"TimeAction\" (user_mail, continent, city, hour, minute, reaction_service_id, reaction_id)"+
		" VALUES($1, $2, $3, $4, $5, $6, $7)", receivedData.Token, receivedData.Continent, receivedData.City, receivedData.Hour, receivedData.Minute, receivedData.ReactionServiceId, 0) // remplacer le reaction_id par le token + la len stp petit con de Edward sale merde mange le caca
	if err != nil {
		c.JSON(http.StatusInsufficientStorage, gin.H{"error": err.Error()})
		return
	}

	defer db.Close(c)
	messageToSend := struct {
		UserToken          string `json:"user_token"`
		ServiceSenderId    int    `json:"service_sender_id"`
		ServiceReceiverId  int    `json:"service_receiver_id"`
		ActionId           int    `json:"action_id"`
		ReactionIdentifyer int    `json:"reaction_identifyer"`
		Message            string `json:"message"`
	}{
		receivedData.Token,
		1,
		2,
		0,
		0,
		receivedData.Message,
	}
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(messageToSend); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	resp, err := http.Post(utils.GetEnvKey("MESSAGE_BROCKER")+"action-message", "application/json", &buf)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"status": resp.Body})
}
