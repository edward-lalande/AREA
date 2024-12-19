package routes

import (
	"fmt"
	"net/http"
	"twilio-service/utils"

	"github.com/gin-gonic/gin"

	area "twilio/Area"
	models "twilio/Models"
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
	row := db.QueryRow(c, "SELECT area_id reaction_type, user_token, phone_number, message FROM \"TwilioReaction\" WHERE area_id = $1", receivedData.ReactionIdentifyer)

	if err := row.Scan(&user.AreaId, &user.ReactionType, &user.UserToken, &user.PhoneNumber, &user.Message); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	defer db.Close(c)

	rep, _ := area.FindReactions(user.ReactionType, models.Reactions{user.PhoneNumber, user.Message})

	c.JSON(rep.StatusCode, gin.H{
		"body": rep.Body,
	})
}
