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

func SendTime(areaId string, action models.TypeTimeAction, c *gin.Context) *http.Response {
	var data models.TimeActionSend
	data.AreaId = areaId
	data.City = action.City
	data.Continent = action.Continent
	data.Hour = action.Hour
	data.Minute = action.Minute

	jsonBody, err := json.Marshal(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil
	}
	resp, err := http.Post(utils.GetEnvKey("TIME_API")+"action", "application/jsons", bytes.NewBuffer(jsonBody))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil
	}
	defer resp.Body.Close()

	return resp
}

func PostTime(c *gin.Context) {
	var body models.TimeDataReceive

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	jsonModels := models.TimeDataToSend{body.Token, body.City, body.Continent, body.Hour, body.Minute, body.ReactionType, body.ReactionServiceId, body.ServerId, body.ChannelId, body.Message}
	jsonBody, err := json.Marshal(jsonModels)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	resp, err := http.Post(utils.GetEnvKey("TIME_API")+body.Routes, "application/jsons", bytes.NewBuffer(jsonBody))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(resp.StatusCode, gin.H{
		"headers": resp.Header,
		"body":    utils.BytesToJson(respBody),
	})
}

func GetTime(c *gin.Context) {
	var body models.DateTimeResponse
	c.ShouldBindJSON(&body)

	resp, err := http.Get(utils.GetEnvKey("TIME_API") + body.Routes)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(resp.StatusCode, gin.H{
		"headers": resp.Header,
		"body":    utils.BytesToJson(respBody),
	})
}
