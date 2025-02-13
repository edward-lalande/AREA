package oauth

import (
	"github/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Github OAUTH2
// @Summary Send the url to redirect to for the OAUTH2 Github
// @Description Send the url to redirect to for the OAUTH2 Github
// @Tags Github OAUTH2
// @Accept json
// @Produce json
// @Success 200 {string} string "the URL to redirect to for the OAUTH2 Github"
// @Router /oauth [get]
func OAuthFront(c *gin.Context) {
	authUrl := "https://github.com/login/oauth/authorize?redirect_uri=" + utils.GetEnvKey("REDIRECT_URI") + "&client_id=" + utils.GetEnvKey("CLIENT_ID")

	c.String(http.StatusOK, authUrl)
}

func AddOAuthFront(c *gin.Context) {
	authUrl := "https://github.com/login/oauth/authorize?redirect_uri=" + utils.GetEnvKey("REDIRECT_URI_ADD") + "&client_id=" + utils.GetEnvKey("CLIENT_ID")

	c.String(http.StatusOK, authUrl)
}
