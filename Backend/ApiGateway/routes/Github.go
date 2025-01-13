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

// GithubOauth2
// @Summary Redirect to Github OAuth2 authorization endpoint
// @Description Initiates the OAuth2 process by redirecting the user to the Github authorization endpoint.
// @Tags Github
// @Produce json
// @Success 200 {string} string "OK"
// @Failure 400 {object} map[string]string "Error message"
// @Router /github/oauth [get]
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

func GithubAddOauth2(c *gin.Context) {

	resp, err := http.Get(utils.GetEnvKey("GITHUB_API") + "add-oauth")

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

// GithubAccessToken
// @Summary Exchange Github OAuth2 authorization code for an access token
// @Description Receives an OAuth2 authorization code and exchanges it for an access token with Github.
// @Tags Github
// @Accept json
// @Produce json
// @Param body body models.OauthCode true "OAuth2 Authorization Code"
// @Success 200 {string} string "Access token response"
// @Failure 400 {object} map[string]string "Error message"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /github/access-token [post]
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

func GithubAddAccessToken(c *gin.Context) {

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

	resp, err := http.Post(utils.GetEnvKey("GITHUB_API")+"add-access-token", "application/json", &buf)
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

// GithubWebhook
// @Summary Handle Github webhook events
// @Description Receives JSON payloads from Github webhook and forwards them to the specified internal service.
// @Tags Github
// @Accept json
// @Produce json
// @Param body body map[string]interface{} true "Github Webhook Payload"
// @Success 200 {object} map[string]string "Webhook forwarded successfully"
// @Failure 400 {object} map[string]string "Invalid JSON"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /github-webhook [post]
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
