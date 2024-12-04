package oauth

import (
	models "discord-service/Models"
	"discord-service/utils"
	"io"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

func GetAccessToken(c *gin.Context) {
	var receivedData models.OauthInformation

	if err := c.ShouldBindJSON(&receivedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	accessTokenUrl := "https://discord.com/api/oauth2/token"
	// userToken, _ := c.GetQuery("token")
	// db := utils.OpenDB(c)
	data := url.Values{}
	data.Set("client_id", utils.GetEnvKey("CLIENT_ID"))
	data.Set("client_secret", utils.GetEnvKey("CLIENT_SECRET"))
	data.Set("redirect_uri", utils.GetEnvKey("REDIRECT_WEB"))
	data.Set("code", receivedData.Code)
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
