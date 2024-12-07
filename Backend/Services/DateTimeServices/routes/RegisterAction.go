package routes

import (
	models "date-time-service/Models"
	"date-time-service/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterAction(c *gin.Context) {
	var dataReceived models.TimeActionReceive
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
		INSERT INTO "TimeAction" (area_id, continent, city, hour, minute)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id;
	`

	_, err := db.Exec(c, query, dataReceived.AreaId, dataReceived.Continent, dataReceived.City, dataReceived.Hour, dataReceived.Minute)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert data into TimeAction: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "TimeAction registered successfully",
		"area_id": dataReceived.AreaId,
	})
	defer db.Close(c)
}
