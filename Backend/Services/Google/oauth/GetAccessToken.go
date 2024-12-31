package oauth

import (
	"bytes"
	"encoding/json"
	"fmt"
	models "google/Models"
	"google/utils"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Google OAUTH2
// @Summary Get
// @Description Send the code received by the frontend to get the Google access-token of the user
// @Tags Google OAUTH2
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

	data, _ := json.Marshal(map[string]string{
		"grant_type":    "authorization_code",
		"code":          receivedData.Code,
		"client_id":     utils.GetEnvKey("CLIENT_ID"),
		"client_secret": utils.GetEnvKey("CLIENT_SECRET"),
		"redirect_uri":  utils.GetEnvKey("REDIRECT_URI"),
	})

	responseBody := bytes.NewBuffer(data)
	rep, err := http.Post("https://oauth2.googleapis.com/token", "application/json", responseBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	respBody, err := io.ReadAll(rep.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("status code: ", rep.StatusCode)

	c.JSON(rep.StatusCode, gin.H{
		"body": utils.BytesToJson(respBody),
	})
}
