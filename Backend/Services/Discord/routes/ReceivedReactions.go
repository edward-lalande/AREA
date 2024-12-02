package routes

import (
	models "discord-service/Models"
	"discord-service/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReceivedReactions(c *gin.Context) {
	var receivedData models.ReactionReceiveData

	db := utils.OpenDB(c)
	if err := c.ShouldBindJSON(&receivedData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	_, err := db.Exec(c, "INSERT INTO \"Reactions\" (service_id, action_id, reaction_identifyer, user_email, message)"+
		" VALUES($1, $2, $3, $4, $5, $6)", receivedData.ServiceId, receivedData.ActionId, receivedData.ReactionIdentifyer, receivedData.UserEmail, receivedData.Message)
	if err != nil {
		c.JSON(http.StatusInsufficientStorage, gin.H{"error": err.Error()})
		return
	}

	db.Close(c)
	c.JSON(http.StatusAccepted, gin.H{"received": receivedData})

}
