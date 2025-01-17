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

func SendAsanaReaction(sendingData models.AsanaReactions, c *gin.Context) *http.Response {

	jsonBody, err := json.Marshal(sendingData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil
	}
	resp, err := http.Post(utils.GetEnvKey("ASANA_API")+"reaction", "application/jsons", bytes.NewBuffer(jsonBody))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil
	}
	defer resp.Body.Close()
	return resp
}

// AsanaOauth2
// @Summary Redirect to Asana OAuth2 authorization endpoint
// @Description Initiates the OAuth2 process by redirecting the user to the Asana authorization endpoint.
// @Tags Asana
// @Produce json
// @Success 200 {string} string "OK"
// @Failure 400 {object} map[string]string "Error message"
// @Router /asana/oauth [get]
func AsanaOauth2(c *gin.Context) {

	resp, err := http.Get(utils.GetEnvKey("ASANA_API") + "oauth")
	fmt.Println("url: ", utils.GetEnvKey("ASANA_API")+"oauth")
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

func AsanaAddOauth2(c *gin.Context) {

	resp, err := http.Get(utils.GetEnvKey("ASANA_API") + "add-oauth")
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

// AsanaAccessToken
// @Summary Exchange Asana OAuth2 authorization code for an access token
// @Description Receives an OAuth2 authorization code and exchanges it for an access token with Asana.
// @Tags Asana
// @Accept json
// @Produce json
// @Param body body models.OauthCode true "OAuth2 Authorization Code"
// @Success 200 {string} string "Access token response"
// @Failure 400 {object} map[string]string "Error message"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /asana/access-token [post]
func AsanaAccessToken(c *gin.Context) {

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

	resp, err := http.Post(utils.GetEnvKey("ASANA_API")+"access-token", "application/json", &buf)
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

func AsanaAddAccessToken(c *gin.Context) {

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

	resp, err := http.Post(utils.GetEnvKey("ASANA_API")+"add-access-token", "application/json", &buf)
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
