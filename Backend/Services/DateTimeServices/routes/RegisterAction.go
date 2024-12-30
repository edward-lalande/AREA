package routes

import (
	models "date-time-service/Models"
	"date-time-service/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Post Time Actions
// @Summary Post an Actions
// @Description Post an time actions, receive by the Message Brocker (handler of communication between services) and register it to him database
// @Tags Actions Date Time services
// @Accept json
// @Produce json
// @Param routes body models.TimeActionReceive true "It contains the Area Id, the location and the time of the Area"
// @Success 200 {object} map[string]string "Response is the Id of the Area"
// @Failure 400 {object} map[string]string "Invalid request"
// @Failure 500 {object} map[string]string "Internal error"
// @Router /actions [post]
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

	_, err := db.Exec(c, "INSERT INTO \"TimeAction\" (area_id, action_type, continent, city, hour, minute)"+
		" VALUES($1, $2, $3, $4, $5, $6)", dataReceived.AreaId, dataReceived.ActionType, dataReceived.Continent, dataReceived.City, dataReceived.Hour, dataReceived.Minute)

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
