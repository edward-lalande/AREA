package routes

import (
	models "api-gateway/Models"
	"api-gateway/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GithubOauth2(c *gin.Context) {

	resp, err := http.Get(utils.GetEnvKey("GITHUB_API") + "oauth")

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

func GithubAccessToken(c *gin.Context) {

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

	resp, err := http.Post(utils.GetEnvKey("GITHUB_API")+"access-token", "application/json", &buf)
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

func GithubWebhook(c *gin.Context) {

	var data map[string]interface{}

	fmt.Println("Webook received!")

	if err := c.ShouldBindJSON(&data); err != nil {
		fmt.Println("Error parsing JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal error"})
		return
	}

	event := c.Request.Header.Get("X-GitHub-Event")

	resp, err := http.Post(utils.GetEnvKey("GITHUB_API")+"webhook/"+event, "application/json", bytes.NewReader(jsonData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, resp.Body)
		return
	}
	defer resp.Body.Close()

	c.JSON(http.StatusOK, gin.H{"message": "Webhook forwarded successfully"})

}

func sendGithub(areaId string, userToken string, c *gin.Context, receivedData models.TypeGithubAction) *http.Response {

	sendingData := struct {
		AreaId     string `json:"area_id"`
		ActionType int    `json:"action_type"`
		UserToken  string `json:"user_token"`
		Pusher     string `json:"pusher"`
		Value      string `json:"value"`
		Number     int    `json:"number"`
	}{
		AreaId:     areaId,
		ActionType: receivedData.ActionType,
		UserToken:  userToken,
		Pusher:     receivedData.Pusher,
		Value:      receivedData.Value,
		Number:     receivedData.Number,
	}

	jsonBody, err := json.Marshal(sendingData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil
	}
	resp, err := http.Post(utils.GetEnvKey("GITHUB_API")+"action", "application/jsons", bytes.NewBuffer(jsonBody))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil
	}
	defer resp.Body.Close()
	return resp

}
