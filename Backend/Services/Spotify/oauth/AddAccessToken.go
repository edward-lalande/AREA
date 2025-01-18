package oauth

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	models "spotify/Models"
	"spotify/utils"

	"github.com/gin-gonic/gin"
)

func AddAccessToken(c *gin.Context) {
	var receivedData models.OauthInformationToken

	if err := c.ShouldBindJSON(&receivedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	accessTokenUrl := "https://accounts.spotify.com/api/token"
	data := url.Values{}
	data.Set("client_id", utils.GetEnvKey("CLIENT_ID"))
	data.Set("client_secret", utils.GetEnvKey("CLIENT_SECRET"))
	data.Set("redirect_uri", utils.GetEnvKey("REDIRECT_URI_ADD"))
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

	if utils.BytesToJson(respBody)["error"] != nil {
		return
	}

	access_token := utils.BytesToJson(respBody)["access_token"]

	if access_token == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(access_token)

	if access_token == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userToken := receivedData.Token
	if userToken == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token is required"})
		return
	}

	fmt.Println("add user token = " + userToken)

	id := utils.ParseToken(userToken)
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token"})
		return
	}

	fmt.Println("add user id  = " + id)

	query := `UPDATE "User" SET spotify_token = $1 WHERE id = $2;`

	db.Exec(c, query, access_token, id)

	c.JSON(rep.StatusCode, "Token registered!")
}
