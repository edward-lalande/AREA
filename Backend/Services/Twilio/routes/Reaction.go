package routes

import (
	"fmt"
	"net/http"
	"twilio-service/utils"

	"github.com/gin-gonic/gin"

	models "twilio/Models"
)

func Reaction(c *gin.Context) {

	var receivedData models.ReactionReceiveData

	db := utils.OpenDB(c)
	if err := c.ShouldBindJSON(&receivedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := db.Exec(c, "INSERT INTO \"TwilioReaction\" (area_id, reaction_type, user_token, phone_number, message)"+
		" VALUES($1, $2, $3, $4, $5)", receivedData.AreaId, receivedData.ReactionType, receivedData.UserToken, receivedData.PhoneNumber, receivedData.Message)

	if err != nil {
		fmt.Println("error:", err.Error())
		c.JSON(http.StatusInsufficientStorage, gin.H{"error": err.Error()})
		return
	}

	defer db.Close(c)

	c.JSON(http.StatusAccepted, gin.H{"Twilio received": receivedData})
}
