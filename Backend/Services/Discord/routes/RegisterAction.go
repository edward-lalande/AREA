package routes

import (
	models "discord-service/Models"
	"discord-service/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterAction(c *gin.Context) {

	var dataReceived models.DiscordActionReceive
	db := utils.OpenDB(c)

	if err := c.ShouldBindJSON(&dataReceived); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to open the database"})
		return
	}

	query := `
		INSERT INTO "DiscordAction" (action_type, channel_id, message_id, area_id, user_token)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id;
	`

	_, err := db.Exec(c, query, dataReceived.Type, dataReceived.AreaId, dataReceived.ChannelId, dataReceived.MessageId, dataReceived.UserToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert data into DiscordAction: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "DiscordAction registered successfully",
		"area_id": dataReceived.AreaId,
	})
	defer db.Close(c)

}
