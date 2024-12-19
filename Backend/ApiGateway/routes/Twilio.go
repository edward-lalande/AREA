package routes

import (
	models "api-gateway/Models"
	"api-gateway/utils"
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func sendTwilioReaction(userToken string, areaId string, c *gin.Context, receivedData models.TypeTwilioReaction) *http.Response {

	sendingData := struct {
		AreaId       string `json:"area_id"`
		UserToken    string `json:"user_token"`
		ReactionType int    `json:"reaction_type"`
		PhoneNumber  string `json:"phone_number"`
		Message      string `json:"message"`
	}{
		AreaId:       areaId,
		UserToken:    userToken,
		ReactionType: receivedData.ReactionType,
		PhoneNumber:  receivedData.PhoneNumber,
		Message:      receivedData.Message,
	}

	jsonBody, err := json.Marshal(sendingData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil
	}
	resp, err := http.Post(utils.GetEnvKey("TWILIO_API")+"reaction", "application/jsons", bytes.NewBuffer(jsonBody))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil
	}
	defer resp.Body.Close()
	return resp

}
