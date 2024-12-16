package routes

import (
	models "discord-service/Models"
	"discord-service/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Discord Services
// @Summary Register an received Reactions
// @Description Register the reactions received by the message brocker with all informations nedded
// @Tags Discord Area
// @Accept json
// @Produce json
// @Param routes body models.ReactionReceiveData true "It must contains the AreaId and the reactions type"
// @Success 200 {object} map[string]string "Response is the received data"
// @Failure 400 {object} map[string]string "Invalid request it contains the error"
// @Failure 500 {object} map[string]string "Internal error it contains the error"
// @Router /reaction [post]
func ReceivedReactions(c *gin.Context) {
	var receivedData models.ReactionReceiveData

	db := utils.OpenDB(c)
	if err := c.ShouldBindJSON(&receivedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := db.Exec(c, "INSERT INTO \"DiscordReactions\" (area_id, reaction_type, user_token, channel_id, message, guild_id)"+
		" VALUES($1, $2, $3, $4, $5, $6)", receivedData.AreaId, receivedData.ReactionType, receivedData.UserToken, receivedData.ChannelID, receivedData.Message, receivedData.GuildID)
	if err != nil {
		fmt.Println("error:", err.Error())
		c.JSON(http.StatusInsufficientStorage, gin.H{"error": err.Error()})
		return
	}

	defer db.Close(c)
	c.JSON(http.StatusAccepted, gin.H{"Discord received": receivedData})
}
