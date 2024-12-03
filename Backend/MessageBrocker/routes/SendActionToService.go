package routes

import (
	"bytes"
	"encoding/json"
	models "message-brocker/Models"
	"message-brocker/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendActionToService(c *gin.Context) {
	var receivedData models.ReceivedActionToReactions
	var services map[int]string = make(map[int]string)

	services[0] = utils.GetEnvKey("USER_API")
	services[1] = utils.GetEnvKey("TIME_API")
	services[2] = utils.GetEnvKey("DISCORD_API")

	if err := c.ShouldBindJSON(&receivedData); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	sendBody := struct {
		ServiceId          int    `json:"service_id"`
		ActionId           int    `json:"action_id"`
		ReactionIdentifyer int    `json:"reaction_identifyer"`
		UserEmail          string `json:"user_email"`
		Message            string `json:"message"`
	}{
		receivedData.ServiceSenderId,
		receivedData.ActionId,
		receivedData.ReactionIdentifyer,
		receivedData.UserToken,
		receivedData.Message,
	}
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(sendBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp, err := http.Post(services[receivedData.ServiceReceiverId]+"create-reactions", "application/json", &buf)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(resp.StatusCode, gin.H{"body": resp.Body})
}
