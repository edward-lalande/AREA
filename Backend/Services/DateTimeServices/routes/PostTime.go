package routes

import (
	"bytes"
	"crypto/rand"
	models "date-time-service/Models"
	"date-time-service/utils"
	"encoding/hex"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GenerateCryptoID() string {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		panic(err)
	}
	return hex.EncodeToString(bytes)
}

func PostTime(c *gin.Context) {
	var receivedData models.DataReceive
	if err := c.ShouldBindJSON(&receivedData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	reactionIdentifyer := GenerateCryptoID()
	db := utils.OpenDB(c)
	_, err := db.Exec(c, "INSERT INTO \"TimeAction\" (user_mail, continent, city, hour, minute, reaction_service_id, reaction_id)"+
		" VALUES($1, $2, $3, $4, $5, $6, $7)", receivedData.Token, receivedData.Continent, receivedData.City, receivedData.Hour, receivedData.Minute, receivedData.ReactionServiceId, reactionIdentifyer)
	if err != nil {
		c.JSON(http.StatusInsufficientStorage, gin.H{"error": err.Error()})
		return
	}

	defer db.Close(c)
	messageToSend := struct {
		UserToken          string `json:"user_token"`
		ServiceSenderId    int    `json:"service_sender_id"`
		ServiceReceiverId  int    `json:"service_receiver_id"`
		ActionId           int    `json:"action_id"`
		ReactionIdentifyer string `json:"area_id"`
		ReactionType       int    `json:"reaction_type"`
		Message            string `json:"message"`
		ServerId           string `json:"server_id"`
		ChannelId          string `json:"channel_id"`
	}{
		receivedData.Token,
		1,
		receivedData.ReactionServiceId,
		0,
		reactionIdentifyer,
		receivedData.ReactionType,
		receivedData.Message,
		receivedData.ServerId,
		receivedData.ChannelId,
	}
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(messageToSend); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	resp, err := http.Post(utils.GetEnvKey("MESSAGE_BROCKER")+"action-message", "application/json", &buf)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(resp.StatusCode, gin.H{"status": resp.Body})
}
