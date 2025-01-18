package oauth

import (
	"context"
	models "discord-service/Models"
	"discord-service/utils"
	"io"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

// Discord OAUTH2
// @Summary Get
// @Description Send the code received by the frontend to get the discord access-token of the user
// @Tags Discord OAUTH2
// @Accept json
// @Produce json
// @Params object models.OauthInformation true "The code must be send as object and the token is not necessary, it can be null"
// @Success 200 {object} map[string]string "the code to redirect to"
// @Router /access-token [post]
func GetAccessToken(c *gin.Context) {
	var receivedData models.OauthInformation

	if err := c.ShouldBindJSON(&receivedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accessTokenUrl := "https://discord.com/api/oauth2/token"
	data := url.Values{}
	data.Set("client_id", utils.GetEnvKey("CLIENT_ID"))
	data.Set("client_secret", utils.GetEnvKey("CLIENT_SECRET"))
	data.Set("redirect_uri", utils.GetEnvKey("REDIRECT_WEB"))
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

	query := `
		INSERT INTO "User" (discord_token)
		VALUES ($1)
		RETURNING id;
	`
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

	var id string

	row := db.QueryRow(context.Background(), query, access_token)
	_ = row.Scan(&id)
	defer db.Close(c)

	token, err := utils.CreateToken(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(rep.StatusCode, gin.H{
		"body": token,
	})
}
