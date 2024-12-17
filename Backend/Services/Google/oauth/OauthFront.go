package oauth

import (
	"google/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Google OAUTH2
// @Summary Send the url to redirect to for the OAUTH2 Google
// @Description Send the url to redirect to for the OAUTH2 Google
// @Tags Google OAUTH2
// @Accept json
// @Produce json
// @Success 200 {string} string "the URL to redirect to for the OAUTH2 Google"
// @Router /oauth2 [get]
func OAuthFront(c *gin.Context) {
	// scope={YOUR_SCOPES_URL_ESCAPED}
	authUrl := "https://accounts.google.com/o/oauth2/auth?client_id=" + utils.GetEnvKey("CLIENT_ID") +
		"&redirect_uri=http://127.0.0.1:8088/callback&response_type=code" +
		"&scope=" + "https://www.googleapis.com/auth/userinfo.profile " +
		"https://www.googleapis.com/auth/userinfo.email " +
		"https://www.googleapis.com/auth/calendar " +
		"https://www.googleapis.com/auth/gmail.readonly " +
		"https://www.googleapis.com/auth/gmail.send " +
		"https://www.googleapis.com/auth/tasks"

	c.String(http.StatusOK, authUrl)
}
