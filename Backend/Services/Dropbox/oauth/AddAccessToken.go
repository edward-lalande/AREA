package oauth

import (
	models "dropbox/Models"
	"dropbox/utils"
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
	accessTokenUrl := "https://api.dropboxapi.com/oauth2/token"
	data := url.Values{}
	data.Set("client_id", utils.GetEnvKey("APP_KEY"))
	data.Set("client_secret", utils.GetEnvKey("APP_SECRET"))
	data.Set("redirect_uri", utils.GetEnvKey("REDIRECT_URI_ADD"))
	data.Set("code", receivedData.Code)
	data.Set("grant_type", "authorization_code")

	rep, _ := http.PostForm(accessTokenUrl, data)
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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

	query := `UPDATE "User" SET dropbox_token = $1 WHERE id = $2;`

	db.Exec(c, query, access_token, id)

	c.JSON(rep.StatusCode, "Token registered!")
}
