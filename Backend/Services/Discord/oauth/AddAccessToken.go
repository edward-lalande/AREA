package oauth

import (
	models "discord-service/Models"
	"discord-service/utils"
	"io"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

func AddAccessToken(c *gin.Context) {
	var receivedData models.OauthInformation

	if err := c.ShouldBindJSON(&receivedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accessTokenUrl := "https://discord.com/api/oauth2/token"
	data := url.Values{}
	data.Set("client_id", utils.GetEnvKey("CLIENT_ID"))
	data.Set("client_secret", utils.GetEnvKey("CLIENT_SECRET"))
	data.Set("redirect_uri", utils.GetEnvKey("REDIRECT_WEB_ADD"))
	data.Set("code", receivedData.Code)
	data.Set("grant_type", "authorization_code")

	rep, err := http.PostForm(accessTokenUrl, data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	respBody, err := io.ReadAll(rep.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	db := utils.OpenDB(c)
	if db == nil {
		return
	}

	if utils.BytesToJson(respBody)["error"] != nil {
		return
	}

	access_token := utils.BytesToJson(respBody)["access_token"]

	if access_token == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No access_token"})
		return
	}

	userToken := receivedData.Token
	if userToken == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token is required"})
		return
	}

	id := utils.ParseToken(userToken)
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token"})
		return
	}

	query := `UPDATE "User" SET discord_token = $1 WHERE id = $2;`

	db.Exec(c, query, access_token, id)

	c.JSON(http.StatusOK, "Token registered!")
}
