package oauth

import (
	"asana/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Asana OAUTH2
// @Summary Send the url to redirect to for the OAUTH2 Asana
// @Description Send the url to redirect to for the OAUTH2 Asana
// @Tags Asana OAUTH2
// @Accept json
// @Produce json
// @Success 200 {string} string "the URL to redirect to for the OAUTH2 Asana"
// @Router /oauth2 [get]
func OAuthFront(c *gin.Context) {
	authUrl := "https://app.asana.com/-/oauth_authorize?client_id=" + utils.GetEnvKey("CLIENT_ID") +
		"&redirect_uri=" + utils.GetEnvKey("REDIRECT_URI") +
		"&response_type=code" +
		"&state=default" +
		"&scope=default" +
		"&grant_type=authorization_code"
	c.String(http.StatusOK, authUrl)
}
