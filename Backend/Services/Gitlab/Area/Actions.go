package area

import (
	models "gitlab/Models"
	"gitlab/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Gitlab Services
// @Summary Register an received Actions
// @Description Register the Actions received by the message brocker with all informations nedded
// @Tags Gitlab Area
// @Accept json
// @Produce json
// @Param routes body models.Actions true "It must contains the AreaId and the reactions type"
// @Success 200 {object} map[string]string "Response is the received data"
// @Failure 400 {object} map[string]string "Invalid request it contains the error"
// @Failure 500 {object} map[string]string "Internal error it contains the error"
// @Router /action [post]
func Actions(c *gin.Context) {
	db := utils.OpenDB(c)
	receivedData := models.Actions{}

	if err := c.ShouldBindJSON(&receivedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	query := `INSERT INTO "GitlabActions" (action_type, area_id) VALUES ($1, $2)`
	_, err := db.Exec(c, query, receivedData.ActionType, receivedData.AreaId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert data"})
		return
	}

	defer db.Close(c)
	c.JSON(http.StatusOK, gin.H{"message": "Data inserted successfully"})
}

// Gitlab Actions
// @Summary send all the Actions
// @Description send all the Actions available on the Gitlab services as an object arrays with the names and the object needed
// @Tags Gitlab Area
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
