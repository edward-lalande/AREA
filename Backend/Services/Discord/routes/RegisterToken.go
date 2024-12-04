package routes

import (
	"context"
	models "discord-service/Models"
	"discord-service/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
