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

func SendGitlabReaction(data models.GitlabReactions, c *gin.Context) *http.Response {
	jsonBody, err := json.Marshal(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil
	}

	resp, err := http.Post(utils.GetEnvKey("GITLAB_API")+"reaction", "application/jsons", bytes.NewBuffer(jsonBody))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil
	}
	defer resp.Body.Close()

	return resp
}

func SendGitlab(areaId string, data models.GitlabAction, c *gin.Context) *http.Response {
	data.AreaId = areaId

	jsonBody, err := json.Marshal(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil
	}

	resp, err := http.Post(utils.GetEnvKey("GITLAB_API")+"action", "application/jsons", bytes.NewBuffer(jsonBody))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil
	}
	defer resp.Body.Close()

	return resp
}

// GitlabWebhook
// @Summary Handle GitLab webhook events
// @Description Receives JSON payloads from GitLab webhook and forwards them to the specified internal service.
// @Tags GitLab
// @Accept json
// @Produce json
// @Param body body map[string]interface{} true "GitLab Webhook Payload"
// @Success 200 {object} map[string]string "Webhook forwarded successfully"
// @Failure 400 {object} map[string]string "Invalid JSON"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /gitlab-webhook [post]
func GitlabWebhook(c *gin.Context) {
	var a map[string]interface{}
	if err := c.ShouldBindJSON(&a); err != nil {
		fmt.Println("Error parsing JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	jsonData, err := json.Marshal(a)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal error"})
		return
	}

	resp, err := http.Post(utils.GetEnvKey("GITLAB_API")+"webhook", "application/json", bytes.NewReader(jsonData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, resp.Body)
		return
	}
	defer resp.Body.Close()

	c.JSON(http.StatusOK, gin.H{"message": "Webhook forwarded successfully"})
}

// GitlabOauth2
// @Summary Redirect to Gitlab OAuth2 authorization endpoint
// @Description Initiates the OAuth2 process by redirecting the user to the Gitlab authorization endpoint.
// @Tags Gitlab
// @Produce json
// @Success 200 {string} string "OK"
// @Failure 400 {object} map[string]string "Error message"
// @Router /gitlab/oauth [get]
func GitlabOauth2(c *gin.Context) {

	resp, err := http.Get(utils.GetEnvKey("GITLAB_API") + "oauth")

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

func GitlabAddOauth2(c *gin.Context) {

	resp, err := http.Get(utils.GetEnvKey("GITLAB_API") + "add-oauth")

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

// GitlabAccessToken
// @Summary Exchange Gitlab OAuth2 authorization code for an access token
// @Description Receives an OAuth2 authorization code and exchanges it for an access token with Gitlab.
// @Tags Gitlab
// @Accept json
// @Produce json
// @Param body body models.OauthCode true "OAuth2 Authorization Code"
// @Success 200 {string} string "Access token response"
// @Failure 400 {object} map[string]string "Error message"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /gitlab/access-token [post]
func GitlabAccessToken(c *gin.Context) {

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

	resp, err := http.Post(utils.GetEnvKey("GITLAB_API")+"access-token", "application/json", &buf)
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

func GitlabAddAccessToken(c *gin.Context) {

	var (
		OauthCode models.OauthCodeToken
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

	resp, err := http.Post(utils.GetEnvKey("GITLAB_API")+"add-access-token", "application/json", &buf)
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
