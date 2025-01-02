package routes

import (
	models "github/Models"
	"github/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getActions(c *gin.Context) {
	b, err := utils.OpenFile("Models/Actions.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	json := utils.BytesToJson(b)

	c.JSON(http.StatusOK, json)
}

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
