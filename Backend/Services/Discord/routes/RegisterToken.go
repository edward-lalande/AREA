package routes

import (
	"context"
	models "discord-service/Models"
	"discord-service/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Discord Services
// @Summary Register a token
// @Description Register the token of the user to an associated token if exists in the user database
// @Tags Discord OAUTH2
// @Accept json
// @Produce json
// @Param routes body models.OauthInformationSignUp true "It must contains the access token of discord and the user token if exists"
// @Success 200 {object} map[string]string "Response of the reactions"
// @Failure 400 {object} map[string]string "Invalid request it contains the error"
// @Failure 500 {object} map[string]string "Internal error it contains the error"
// @Router /register [post]
func RegisterToken(c *gin.Context) {
	var oauthInfo models.OauthInformationSignUp

	if err := c.ShouldBindJSON(&oauthInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	db := utils.OpenDB(c)
	if db == nil {
		return
	}
	defer db.Close(context.Background())

	query := `
		INSERT INTO "DiscordUser" (user_token, access_token)
		VALUES ($1, $2)
		RETURNING id;
	`

	var insertedID int
	err := db.QueryRow(context.Background(), query, oauthInfo.UserToken, oauthInfo.AccessToken).Scan(&insertedID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert user", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User registered successfully",
		"user_id": insertedID,
	})
}
