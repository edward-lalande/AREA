package oauth

import (
	models "dropbox/Models"
	"dropbox/utils"
	"io"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

// Dropbox OAUTH2
// @Summary Get
// @Description Send the code received by the frontend to get the Dropbox access-token of the user
// @Tags Dropbox OAUTH2
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
	accessTokenUrl := "https://api.dropboxapi.com/oauth2/token"
	data := url.Values{}
	data.Set("client_id", utils.GetEnvKey("APP_KEY"))
	data.Set("client_secret", utils.GetEnvKey("APP_SECRET"))
	data.Set("redirect_uri", utils.GetEnvKey("REDIRECT_URI"))
	data.Set("code", receivedData.Code)
	data.Set("grant_type", "authorization_code")

	rep, _ := http.PostForm(accessTokenUrl, data)
	respBody, err := io.ReadAll(rep.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	query := `
		INSERT INTO "User" (dropbox_token)
		VALUES ($1)
		RETURNING id;
	`

	db := utils.OpenDB(c)
	if db == nil {
		return
	}

	var id string
	db.QueryRow(c, query, utils.BytesToJson(respBody)).Scan(&id)

	token, err := utils.CreateToken(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(rep.StatusCode, gin.H{
		"body": token,
	})
}
