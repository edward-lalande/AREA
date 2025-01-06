package controllers

import (
	"net/http"
	models "ticket-master/Models"
	"ticket-master/utils"

	"github.com/gin-gonic/gin"
)

// Ticket Master Actions
// @Summary send all the Actions
// @Description send all the Actions available on the Ticket Master services as an object arrays with the names and the object needed
// @Tags Ticket Master Area
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string "Response is the received data"
// @Failure 400 {object} map[string]string "Invalid request it contains the error"
// @Failure 500 {object} map[string]string "Internal error it contains the error"
// @Router /actions [get]
func GetActions(c *gin.Context) {
	b, err := utils.OpenFile("Models/Actions.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	json := utils.BytesToJson(b)
	c.JSON(http.StatusOK, json)
}

// TIcketMaster Services
// @Summary Register an received Actions
// @Description Register the Actions received by the message brocker with all informations nedded
// @Tags TIcketMaster Area
// @Accept json
// @Produce json
// @Param routes body models.Action true "It must contains the AreaId and the reactions type"
// @Success 200 {object} map[string]string "Response is the received data"
// @Failure 400 {object} map[string]string "Invalid request it contains the error"
// @Failure 500 {object} map[string]string "Internal error it contains the error"
// @Router /action [post]
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
