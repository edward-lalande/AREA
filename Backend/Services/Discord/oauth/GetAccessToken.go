package oauth

import (
	"discord-service/utils"
	"io"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

func GetAccessToken(c *gin.Context) {
	accessTokenUrl := "https://discord.com/api/oauth2/token"
	code, _ := c.GetQuery("code")
	data := url.Values{}

	data.Set("client_id", utils.GetEnvKey("CLIENT_ID"))
	data.Set("client_secret", utils.GetEnvKey("CLIENT_SECRET"))
	data.Set("redirect_uri", utils.GetEnvKey("REDIRECT_WEB"))
	data.Set("code", code)
	data.Set("grant_type", "authorization_code")

	rep, _ := http.PostForm(accessTokenUrl, data)
	respBody, err := io.ReadAll(rep.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(rep.StatusCode, gin.H{
		"body": utils.BytesToJson(respBody),
	})
}
