package oauth

import (
	"miro/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Miro OAUTH2
// @Summary Send the url to redirect to for the OAUTH2 Miro
// @Description Send the url to redirect to for the OAUTH2 Miro
// @Tags Miro OAUTH2
// @Accept json
// @Produce json
// @Success 200 {string} string "the URL to redirect to for the OAUTH2 Miro"
// @Router /oauth [get]
func OAuthFront(c *gin.Context) {
	authUrl := "https://miro.com/oauth/authorize?response_type=code" +
		"&client_id=" + utils.GetEnvKey("CLIENT_ID") +
		"&redirect_uri=" + utils.GetEnvKey("CLIENT_SECRET")
	c.String(http.StatusOK, authUrl)
}
