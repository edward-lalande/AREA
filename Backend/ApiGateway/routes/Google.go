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

func SendGoogleReactions(areaId string, data models.GoogleReaction, c *gin.Context) *http.Response {
	data.AreaId = areaId

	jsonBody, err := json.Marshal(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil
	}

	resp, err := http.Post(utils.GetEnvKey("GOOGLE_API")+"reaction", "application/jsons", bytes.NewBuffer(jsonBody))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil
	}
	defer resp.Body.Close()

	return resp
}

func SendGoogle(areaId string, data models.GoogleAction, c *gin.Context) *http.Response {
	data.AreaId = areaId

	jsonBody, err := json.Marshal(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil
	}

	resp, err := http.Post(utils.GetEnvKey("GOOGLE_API")+"action", "application/jsons", bytes.NewBuffer(jsonBody))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil
	}
	defer resp.Body.Close()

	return resp
}

// GoogleOauth2
// @Summary Redirect to Google OAuth2 authorization endpoint
// @Description Initiates the OAuth2 process by redirecting the user to the Google authorization endpoint.
// @Tags Google
// @Produce json
// @Success 200 {string} string "OK"
// @Failure 400 {object} map[string]string "Error message"
// @Router /google/oauth [get]
func GoogleOauth2(c *gin.Context) {

	resp, err := http.Get(utils.GetEnvKey("GOOGLE_API") + "oauth")

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

func GoogleAddOauth2(c *gin.Context) {

	resp, err := http.Get(utils.GetEnvKey("GOOGLE_API") + "add-oauth")

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

// GoogleAccessToken
// @Summary Exchange Google OAuth2 authorization code for an access token
// @Description Receives an OAuth2 authorization code and exchanges it for an access token with Google.
// @Tags Google
// @Accept json
// @Produce json
// @Param body body models.OauthCode true "OAuth2 Authorization Code"
// @Success 200 {string} string "Access token response"
// @Failure 400 {object} map[string]string "Error message"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /google/access-token [post]
func GoogleAccessToken(c *gin.Context) {

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

	resp, err := http.Post(utils.GetEnvKey("GOOGLE_API")+"access-token", "application/json", &buf)
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

func GoogleAddAccessToken(c *gin.Context) {

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

	resp, err := http.Post(utils.GetEnvKey("GOOGLE_API")+"add-access-token", "application/json", &buf)
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
