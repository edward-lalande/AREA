package controllers

import (
	"net/http"
	models "ticket-master/Models"
	"ticket-master/utils"

	"github.com/gin-gonic/gin"
)

func GetActions(c *gin.Context) {
	b, err := utils.OpenFile("Models/Actions.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	json := utils.BytesToJson(b)
	c.JSON(http.StatusOK, json)
}

func StoreActions(c *gin.Context) {
	receivedData := models.Action{}

	if err := c.ShouldBindJSON(&receivedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data", "details": err.Error()})
		return
	}

	if receivedData.AreaID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "AreaID are required fields"})
		return
	}

	db := utils.OpenDB(c)
	defer db.Close(c)

	query := `
		INSERT INTO "TicketMasterActions" (area_id, action_type, name, venue, city, nb_events)
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err := db.Exec(c, query,
		receivedData.AreaID,
		receivedData.ActionType,
		receivedData.Name,
		receivedData.Venue,
		receivedData.City,
		0,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store action", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Action stored successfully"})
}
