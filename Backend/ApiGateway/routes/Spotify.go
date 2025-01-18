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

func SendSpotifyReactions(data models.SpotifyReactions, c *gin.Context) *http.Response {
	jsonBody, err := json.Marshal(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil
	}

	resp, err := http.Post(utils.GetEnvKey("SPOTIFY_API")+"reaction", "application/jsons", bytes.NewBuffer(jsonBody))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil
	}
	defer resp.Body.Close()

	return resp
}

func SendSpotifyActions(data models.SpotifyActions, c *gin.Context) *http.Response {
	jsonBody, err := json.Marshal(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil
	}

	resp, err := http.Post(utils.GetEnvKey("SPOTIFY_API")+"action", "application/jsons", bytes.NewBuffer(jsonBody))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil
	}
	defer resp.Body.Close()

	return resp
}

// SpotifyOauth2
// @Summary Redirect to Spotify OAuth2 authorization endpoint
// @Description Initiates the OAuth2 process by redirecting the user to the Spotify authorization endpoint.
// @Tags Spotify
// @Produce json
// @Success 200 {string} string "OK"
// @Failure 400 {object} map[string]string "Error message"
// @Router /spotify/oauth [get]
func SpotifyOauth2(c *gin.Context) {

	resp, err := http.Get(utils.GetEnvKey("SPOTIFY_API") + "oauth")

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

func SpotifyAddOauth2(c *gin.Context) {

	resp, err := http.Get(utils.GetEnvKey("SPOTIFY_API") + "add-oauth")

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

// SpotifyAccessToken
// @Summary Exchange Spotify OAuth2 authorization code for an access token
// @Description Receives an OAuth2 authorization code and exchanges it for an access token with Spotify.
// @Tags Spotify
// @Accept json
// @Produce json
// @Param body body models.OauthCode true "OAuth2 Authorization Code"
// @Success 200 {string} string "Access token response"
// @Failure 400 {object} map[string]string "Error message"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /spotify/access-token [post]
func SpotifyAccessToken(c *gin.Context) {

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

	resp, err := http.Post(utils.GetEnvKey("SPOTIFY_API")+"access-token", "application/json", &buf)
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

func SpotifyAddAccessToken(c *gin.Context) {

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

	resp, err := http.Post(utils.GetEnvKey("SPOTIFY_API")+"add-access-token", "application/json", &buf)
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
