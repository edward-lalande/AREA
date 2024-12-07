package routes

import (
	models "api-gateway/Models"
	"api-gateway/utils"
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendMessageDiscordReaction(userToken string, areaId string, c *gin.Context, receivedData models.TypeDiscordReaction) *http.Response {
	sendingData := struct {
		AreaId       string `json:"area_id"`
		UserToken    string `json:"user_token"`
		ReactionType int    `json:"reaction_type"`
		ChannelID    string `json:"channel_id"`
		Message      string `json:"message"`
	}{
		AreaId:       areaId,
		UserToken:    userToken,
		ReactionType: receivedData.ReactionType,
		ChannelID:    receivedData.ChannelID,
		Message:      receivedData.Message,
	}

	jsonBody, err := json.Marshal(sendingData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil
	}
	resp, err := http.Post(utils.GetEnvKey("DISCORD_API")+"reaction", "application/jsons", bytes.NewBuffer(jsonBody))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil
	}
	defer resp.Body.Close()
	return resp
}

func DiscordGet(c *gin.Context) {
	var data models.DiscordGet

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := http.Get(utils.GetEnvKey("DISCORD_API") + data.Routes)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	for key, values := range resp.Header {
		for _, value := range values {
			c.Header(key, value)
		}
	}

	c.Status(resp.StatusCode)
	io.Copy(c.Writer, resp.Body)
}

func DiscordPost(c *gin.Context) {
	var (
		data        models.DiscordPost
		sendingData models.DiscordPostOatuh
	)

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sendingData.Code = data.Code
	sendingData.Token = data.Token

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(sendingData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp, err := http.Post(utils.GetEnvKey("DISCORD_API")+data.Routes, "application/json", &buf)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	for key, values := range resp.Header {
		for _, value := range values {
			c.Header(key, value)
		}
	}

	c.Status(resp.StatusCode)
	io.Copy(c.Writer, resp.Body)
}
