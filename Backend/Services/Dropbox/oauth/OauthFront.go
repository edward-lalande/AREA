package oauth

import (
	"dropbox/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Dropbox OAUTH2
// @Summary Send the url to redirect to for the OAUTH2 Dropbox
// @Description Send the url to redirect to for the OAUTH2 Dropbox
// @Tags Dropbox OAUTH2
// @Accept json
// @Produce json
// @Success 200 {string} string "the URL to redirect to for the OAUTH2 Dropbox"
// @Router /oauth2 [get]
func OAuthFront(c *gin.Context) {
	authUrl := "https://www.dropbox.com/oauth2/authorize?&client_id=" + utils.GetEnvKey("APP_KEY") +
		"&redirect_uri=" + utils.GetEnvKey("REDIRECT_URI") +
		"&response_type=code"

	c.String(http.StatusOK, authUrl)
}
