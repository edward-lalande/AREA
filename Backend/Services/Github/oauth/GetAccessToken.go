package oauth

import (
	"fmt"
	models "github/Models"
	"github/utils"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

// Github OAUTH2
// @Summary Get
// @Description Send the code received by the frontend to get the Github access-token of the user
// @Tags Github OAUTH2
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
	fmt.Println("code => " + receivedData.Code)
	accessTokenUrl := "https://github.com/login/oauth/access_token"
	data := url.Values{}
	data.Set("client_id", utils.GetEnvKey("CLIENT_ID"))
	data.Set("client_secret", utils.GetEnvKey("CLIENT_SECRET"))
	data.Set("redirect_uri", utils.GetEnvKey("REDIRECT_URI"))
	data.Set("code", receivedData.Code)
	data.Set("grant_type", "authorization_code")

	rep, _ := http.PostForm(accessTokenUrl, data)
	respBody, err := io.ReadAll(rep.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if rep.StatusCode < 400 {
		arr := strings.Split(string(respBody), "&")
		token := strings.Split(arr[0], "=")
		c.JSON(rep.StatusCode, gin.H{
			"body": token[1],
		})
		return
	}
	c.JSON(400, "Error")
}
