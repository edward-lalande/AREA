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
	data.ActionType = action.ActionType
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
