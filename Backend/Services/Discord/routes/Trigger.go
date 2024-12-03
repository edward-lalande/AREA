package routes

import (
	models "discord-service/Models"
	"discord-service/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Trigger(c *gin.Context) {
	var (
		receivedData models.TriggerdModels
		user         models.TriggerdUserModel
	)

	if err := c.ShouldBindJSON(&receivedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := utils.OpenDB(c)

	row := db.QueryRow(c, "SELECT user_email, message FROM \"DiscordReactions\" WHERE reaction_identifyer = $1", receivedData.ReactionIdentifyer)

	if err := row.Scan(&user.UserEmail, &user.Message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer db.Close(c)

	fmt.Printf("id: (%d), Message: (%s), User: (%s)\n", receivedData.ReactionIdentifyer, user.UserEmail, user.Message)
	c.JSON(http.StatusOK, gin.H{
		"id":      receivedData.ReactionIdentifyer,
		"user":    user.UserEmail,
		"message": user.Message,
	})
}
