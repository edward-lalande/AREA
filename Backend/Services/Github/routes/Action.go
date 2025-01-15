package routes

import (
	models "github/Models"
	"github/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Github Actions
// @Summary send all the Actions
// @Description send all the Actions available on the Github services as an object arrays with the names and the object needed
// @Tags Github Area
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string "Response is the received data"
// @Failure 400 {object} map[string]string "Invalid request it contains the error"
// @Failure 500 {object} map[string]string "Internal error it contains the error"
// @Router /actions [get]
func getActions(c *gin.Context) {
	b, err := utils.OpenFile(models.ActionsModelsPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	json := utils.BytesToJson(b)

	c.JSON(http.StatusOK, json)
}

// Github Services
// @Summary Register an received Actions
// @Description Register the Actions received by the message brocker with all informations nedded
// @Tags Github Area
// @Accept json
// @Produce json
// @Param routes body models.GithubAction true "It must contains the AreaId and the reactions type"
// @Success 200 {object} map[string]string "Response is the received data"
// @Failure 400 {object} map[string]string "Invalid request it contains the error"
// @Failure 500 {object} map[string]string "Internal error it contains the error"
// @Router /action [post]
func createAction(c *gin.Context) {

	var dataReceived models.GithubAction
	db := utils.OpenDB(c)

	if err := c.ShouldBindJSON(&dataReceived); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to open the database"})
		return
	}

	_, err := db.Exec(c, "INSERT INTO \"GithubActions\" (area_id, action_type, user_token, pusher, value, number)"+
		" VALUES($1, $2, $3, $4, $5, $6)", dataReceived.AreaId, dataReceived.ActionType, dataReceived.UserToken, dataReceived.Pusher, dataReceived.Value, dataReceived.Number)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert data into GithubActions: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "GithubActions registered successfully",
		"area_id": dataReceived.AreaId,
	})
	defer db.Close(c)

}
