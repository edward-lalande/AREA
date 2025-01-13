package oauth

import (
	"io"
	models "miro/Models"
	"miro/utils"
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
	accessTokenUrl := "https://api.miro.com/v1/oauth/token"
	data := url.Values{}
	data.Set("client_id", utils.GetEnvKey("CLIENT_ID"))
	data.Set("client_secret", utils.GetEnvKey("CLIENT_SECRET"))
	data.Set("redirect_uri", utils.GetEnvKey("REDIRECT_URI"))
	data.Set("code", receivedData.Code)
	data.Set("grant_type", "authorization_code")

	rep, err := http.PostForm(accessTokenUrl, data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if rep.StatusCode > 200 {
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

	access_token := utils.BytesToJson(respBody)

	if access_token == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userToken := c.GetHeader("token")
	if userToken == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token is required"})
		return
	}

	id := utils.ParseToken(userToken)
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token"})
		return
	}

	query := `UPDATE "User" SET miro_token = $1 WHERE id = $2;`

	db.Exec(c, query, access_token, id)

	c.JSON(rep.StatusCode, "Token registered!")
}
