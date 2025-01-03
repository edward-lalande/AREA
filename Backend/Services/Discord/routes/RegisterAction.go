package routes

import (
	models "discord-service/Models"
	"discord-service/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Discord Services
// @Summary Register an received Actions
// @Description Register the Actions received by the message brocker with all informations nedded
// @Tags Discord Area
// @Accept json
// @Produce json
// @Param routes body models.DiscordActionReceive true "It must contains the AreaId and the reactions type"
// @Success 200 {object} map[string]string "Response is the received data"
// @Failure 400 {object} map[string]string "Invalid request it contains the error"
// @Failure 500 {object} map[string]string "Internal error it contains the error"
// @Router /reaction [post]
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
