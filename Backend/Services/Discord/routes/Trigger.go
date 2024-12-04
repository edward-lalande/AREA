package routes

import (
	area "discord-service/Area"
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
	row := db.QueryRow(c, "SELECT user_token, reaction_type, message, channel_id FROM \"DiscordReactions\" WHERE reaction_identifyer = $1", receivedData.ReactionIdentifyer)

	if err := row.Scan(&user.UserEmail, &user.ReactionType, &user.Message, &user.Channel); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer db.Close(c)

	rep, _ := area.FindReactions(user.ReactionType, models.Reactions{user.Message, user.Channel})
	c.JSON(rep.StatusCode, gin.H{
		"body": rep.Body,
	})
}
