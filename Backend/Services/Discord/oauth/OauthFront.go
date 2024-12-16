package oauth

import (
	"discord-service/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Discord OAUTH2
// @Summary Send the url to redirect to for the OAUTH2 discord
// @Description Send the url to redirect to for the OAUTH2 discord
// @Tags Discord OAUTH2
// @Accept json
// @Produce json
// @Success 200 {string} string "the URL to redirect to for the OAUTH2 discord"
// @Router /oauth2 [get]
func OAuthFront(c *gin.Context) {
	permissions := "2048%20" + "16%20" + "8%20"
	authUrl := "https://discord.com/oauth2/authorize?&client_id=" + utils.GetEnvKey("CLIENT_ID") +
		"&redirect_uri=" + utils.GetEnvKey("REDIRECT_WEB") +
		"&response_type=code" +
		"&scope=identify%20email%20guilds%20bot" +
		"&permissions=" + permissions

	c.String(http.StatusOK, authUrl)
}
