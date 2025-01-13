package oauth

import (
	"bytes"
	"encoding/json"
	models "google/Models"
	"google/utils"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddAccessToken(c *gin.Context) {
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

	query := `UPDATE "User" SET google_token = $1 WHERE id = $2;`

	db.Exec(c, query, access_token, id)

	c.JSON(rep.StatusCode, "Token registered!")
}
