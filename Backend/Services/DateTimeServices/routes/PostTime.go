package routes

import (
	models "date-time-service/Models"
	"date-time-service/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostTime(c *gin.Context) {
	var receivedData models.DataReceive
	db := utils.OpenDB(c)
	if err := c.ShouldBindJSON(&receivedData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	_, err := db.Exec(c, "INSERT INTO \"Action\" (user_mail, continent, city, hour, minute)"+
		" VALUES($1, $2, $3, $4, $5)", receivedData.Token, receivedData.Continent, receivedData.City, receivedData.Hour, receivedData.Minute)
	if err != nil {
		c.JSON(http.StatusInsufficientStorage, gin.H{"error": err.Error()})
		return
	}

	db.Close(c)
	c.JSON(http.StatusAccepted, gin.H{"status": "0"})
}
