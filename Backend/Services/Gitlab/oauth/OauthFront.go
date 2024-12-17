package oauth

import (
	"gitlab/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Gitlab OAUTH2
// @Summary Send the url to redirect to for the OAUTH2 Gitlab
// @Description Send the url to redirect to for the OAUTH2 Gitlab
// @Tags Gitlab OAUTH2
// @Accept json
// @Produce json
// @Success 200 {string} string "the URL to redirect to for the OAUTH2 Gitlab"
// @Router /oauth2 [get]
func OAuthFront(c *gin.Context) {
	authUrl := "https://gitlab.com/oauth/authorize?client_id=" + utils.GetEnvKey("CLIENT_ID") +
		"&client_secret=" + utils.GetEnvKey("CLIENT_SECRET") +
		"&redirect_uri=http://127.0.0.1:8087/callback&response_type=code" +
		"&scope=" + "api " + "read_api " + "read_user " + "create_runner " + "manage_runner " + "admin_mode " + ""
	c.String(http.StatusOK, authUrl)
}
