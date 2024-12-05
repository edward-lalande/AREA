package routes

import (
	"bytes"
	"encoding/json"
	models "message-brocker/Models"
	"message-brocker/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Trigger(c *gin.Context) {
	var (
		receivedData models.TriggerModelGateway
		services     map[int]string = make(map[int]string)
		sendBody     models.TriggerdModelsSending
	)
	services[0] = utils.GetEnvKey("USER_API")
	services[1] = utils.GetEnvKey("TIME_API")
	services[2] = utils.GetEnvKey("DISCORD_API")

	if err := c.ShouldBindJSON(&receivedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	sendBody.ReactionIdentifyer = receivedData.ReactionId
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(sendBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp, err := http.Post(services[receivedData.ReactionServiceId]+"trigger", "application/json", &buf)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(resp.StatusCode, gin.H{"body": resp.Body})
}
