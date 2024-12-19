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
	services[3] = utils.GetEnvKey("BLUESKY_API")
	services[4] = utils.GetEnvKey("GITHUB_API")
	services[5] = utils.GetEnvKey("GITLAB_API")
	services[6] = utils.GetEnvKey("GOOGLE_API")
	services[7] = utils.GetEnvKey("METEO_API")
	services[8] = utils.GetEnvKey("MIRO_API")
	services[9] = utils.GetEnvKey("SPOTIFY_API")
	services[10] = utils.GetEnvKey("STEAM_API")
	services[11] = utils.GetEnvKey("TICKET_MASTER_API")
	services[12] = utils.GetEnvKey("TWILIO_API")
	services[13] = utils.GetEnvKey("UBER_API")

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
		PhoneNumber        string `json:"phone_number"`
	}{
		receivedData.ReactionIdentifyer,
		receivedData.ReactionType,
		receivedData.UserToken,
		receivedData.Message,
		receivedData.ChannelId,
		receivedData.GuildId,
		receivedData.PhoneNumber,
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
