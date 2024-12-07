package routes

import (
	models "discord-service/Models"
	"discord-service/utils"
	"fmt"
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

	_, err := db.Exec(c, "INSERT INTO \"DiscordReactions\" (area_id, reaction_type, user_token, channel_id, message)"+
		" VALUES($1, $2, $3, $4, $5)", receivedData.AreaId, receivedData.ReactionType, receivedData.UserToken, receivedData.ChannelID, receivedData.Message)
	if err != nil {
		fmt.Println("error:", err.Error())
		c.JSON(http.StatusInsufficientStorage, gin.H{"error": err.Error()})
		return
	}

	defer db.Close(c)
	c.JSON(http.StatusAccepted, gin.H{"Discord received": receivedData})
}
