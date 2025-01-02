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
	receivedData := models.ReceivedActionToReactions{}
	services := map[int]string{
		0:  utils.GetEnvKey("USER_API"),
		1:  utils.GetEnvKey("TIME_API"),
		2:  utils.GetEnvKey("DISCORD_API"),
		3:  utils.GetEnvKey("DROPBOX_API"),
		4:  utils.GetEnvKey("GITHUB_API"),
		5:  utils.GetEnvKey("GITLAB_API"),
		6:  utils.GetEnvKey("GOOGLE_API"),
		7:  utils.GetEnvKey("METEO_API"),
		9:  utils.GetEnvKey("SPOTIFY_API"),
		10: utils.GetEnvKey("ASANA_API"),
		11: utils.GetEnvKey("TICKET_MASTER_API"),
		12: utils.GetEnvKey("TWILIO_API"),
	}

	if err := c.ShouldBindJSON(&receivedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sendBody := struct {
		ReactionIdentifyer string `json:"area_id"`
		ReactionType       int    `json:"reaction_type"`
		UserEmail          string `json:"user_email"`
		Message            string `json:"message"`
		ChannelId          string `json:"channel_id"`
		GuildId            string `json:"guild_id"`
	}{
		receivedData.ReactionIdentifyer,
		receivedData.ReactionType,
		receivedData.UserToken,
		receivedData.Message,
		receivedData.ChannelId,
		receivedData.GuildId,
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
