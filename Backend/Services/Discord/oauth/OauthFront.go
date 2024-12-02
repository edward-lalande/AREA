package oauth

import (
	"discord-service/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func OAuthFront(c *gin.Context) {
	permissions := "2048%20" + "16%20"
	authUrl := "https://discord.com/oauth2/authorize?&client_id=" + utils.GetEnvKey("CLIENT_ID") +
		"&redirect_uri=" + utils.GetEnvKey("REDIRECT_WEB") +
		"&response_type=code" +
		"&scope=identify%20email%20guilds%20bot" +
		"&permissions=" + permissions

	c.String(http.StatusOK, authUrl)
}
