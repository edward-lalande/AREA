package oauth

import (
	"net/http"
	"spotify/utils"

	"github.com/gin-gonic/gin"
)

// Spotify OAUTH2
// @Summary Send the url to redirect to for the OAUTH2 spotify
// @Description Send the url to redirect to for the OAUTH2 spotify
// @Tags spotify OAUTH2
// @Accept json
// @Produce json
// @Success 200 {string} string "the URL to redirect to for the OAUTH2 spotify"
// @Router /oauth2 [get]
func OAuthFront(c *gin.Context) {
	authUrl := "https://accounts.spotify.com/authorize?&client_id=" + utils.GetEnvKey("CLIENT_ID") +
		"&redirect_uri=" + utils.GetEnvKey("REDIRECT_URI") +
		"&response_type=code" +
		"&scope=user-read-private%20user-read-email"
	c.String(http.StatusOK, authUrl)
}
