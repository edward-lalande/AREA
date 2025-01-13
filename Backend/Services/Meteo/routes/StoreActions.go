package routes

import (
	models "meteo/Models"
	"meteo/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Post Meteo Actions
// @Summary Post an Actions
// @Description Post an Meteo actions, receive by the Message Brocker (handler of communication between services) and register it to him database
// @Tags Actions Meteo services
// @Accept json
// @Produce json
// @Param routes body models.MeteoActions true "It contains the Area Id, the location and the Meteo of the Area"
// @Success 200 {object} map[string]string "Response is the Id of the Area"
// @Failure 400 {object} map[string]string "Invalid request"
// @Failure 500 {object} map[string]string "Internal error"
// @Router /actions [post]
func StoreActions(c *gin.Context) {
	receivedData := models.MeteoActions{}
	db := utils.OpenDB(c)
	defer db.Close(c)

	if err := c.ShouldBindJSON(&receivedData); err != nil {
		c.JSON(http.StatusBadRequest, "Invalid Request")
		return
	}

	query := `
		INSERT INTO "MeteoActions" (area_id, action_type, latitude, longitude, value)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id;
	`

	_, err := db.Exec(c, query, receivedData.AreaId, receivedData.ActionType, receivedData.Latitude, receivedData.Longitude, receivedData.Value)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert data into MeteoActions", "details": err.Error()})
		return
	}
}

// Get Actions of Meteo Services
// @Summary Get Actions from Meteo Services
// @Description Get Actions from Meteo Services
// @Tags Actions Meteo Services
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string "Reactions name with parameters of it as object"
// @Router /actions [get]
func GetActions(c *gin.Context) {
	b, err := utils.OpenFile(models.GetActionsPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	json := utils.BytesToJson(b)
	c.JSON(http.StatusOK, json)
}
