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

func sendDiscordAction(userToken string, areaId string, c *gin.Context, receivedData models.TypeDiscordAction) *http.Response {

	sendingData := struct {
		AreaId     string `json:"area_id"`
		UserToken  string `json:"user_token"`
		ActionType int    `json:"action_type"`
		ChannelID  string `json:"channel_id"`
		MessageID  string `json:"message_id"`
	}{
		AreaId:     areaId,
		UserToken:  userToken,
		ActionType: receivedData.ActionType,
		ChannelID:  receivedData.ChannelId,
		MessageID:  receivedData.MessageId,
	}

	jsonBody, err := json.Marshal(sendingData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil
	}
	resp, err := http.Post(utils.GetEnvKey("DISCORD_API")+"action", "application/jsons", bytes.NewBuffer(jsonBody))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil
	}
	defer resp.Body.Close()
	return resp

}

func SendMessageDiscordReaction(userToken string, areaId string, c *gin.Context, receivedData models.TypeDiscordReaction) *http.Response {
	sendingData := struct {
		AreaId       string `json:"area_id"`
		UserToken    string `json:"user_token"`
		ReactionType int    `json:"reaction_type"`
		ChannelID    string `json:"channel_id"`
		Message      string `json:"message"`
		GuildID      string `json:"guild_id"`
	}{
		AreaId:       areaId,
		UserToken:    userToken,
		ReactionType: receivedData.ReactionType,
		ChannelID:    receivedData.ChannelID,
		Message:      receivedData.Message,
		GuildID:      receivedData.GuildID,
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

// Get Discord services
// @Summary Get Data from discord services
// @Description Get data from discord like ping, access-token...
// @Tags Discord api-gateway
// @Accept json
// @Produce json
// @Param routes body models.DiscordGet true "routes you would like to access to Discord"
// @Success 200 {object} map[string]string "Response of Discord"
// @Failure 400 {object} map[string]string "Invalid request"
// @Failure 500 {object} map[string]string "Internal error"
// @Router /discord [get]
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

// Post Discord services
// @Summary Post Data to discord services
// @Description Post data to discord for oauth
// @Tags Discord api-gateway
// @Accept json
// @Produce json
// @Param routes body models.DiscordPost true "routes you would like to access to Discord, code of the user and token of him if already exists"
// @Success 200 {object} map[string]string "Response of Discord"
// @Failure 400 {object} map[string]string "Invalid request"
// @Failure 500 {object} map[string]string "Internal error"
// @Router /discord [post]
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

func DiscordOauth2(c *gin.Context) {

	resp, err := http.Get(utils.GetEnvKey("DISCORD_API") + "oauth")

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

func DiscordAccessToken(c *gin.Context) {

	var (
		OauthCode models.OauthCode
	)

	if err := c.ShouldBindJSON(&OauthCode); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(OauthCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp, err := http.Post(utils.GetEnvKey("DISCORD_API")+"access-token", "application/json", &buf)
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
