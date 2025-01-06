package oauth

import (
	"io"
	models "miro/Models"
	"miro/utils"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

// Miro OAUTH2
// @Summary Get
// @Description Send the code received by the frontend to get the Miro access-token of the user
// @Tags Miro OAUTH2
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
	accessTokenUrl := "https://api.miro.com/v1/oauth/token"
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
