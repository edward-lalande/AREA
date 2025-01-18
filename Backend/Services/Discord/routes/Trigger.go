package routes

import (
	area "discord-service/Area"
	models "discord-service/Models"
	"discord-service/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Discord Services
// @Summary Trigger an Area
// @Description Actions triggerd the reactions and call the trigger route
// @Tags Discord trigger
// @Accept json
// @Produce json
// @Param routes body models.TriggerdModels true "It contains the Area Id to the reactions"
// @Success 200 {object} map[string]string "Response of the reactions"
// @Failure 400 {object} map[string]string "Invalid request it contains the error"
// @Failure 500 {object} map[string]string "Internal error it contains the error"
// @Router /trigger [post]
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
	row := db.QueryRow(c, "SELECT reaction_type, message, channel_id, guild_id FROM \"DiscordReactions\" WHERE area_id = $1", receivedData.AreaId)

	if err := row.Scan(&user.ReactionType, &user.Message, &user.Channel, &user.Guild); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer db.Close(c)

	rep, _ := area.FindReactions(user.ReactionType, models.Reactions{Message: user.Message, ChannelId: user.Channel, GuildId: user.Guild})

	c.JSON(rep.StatusCode, gin.H{
		"body": rep.Body,
	})
}
