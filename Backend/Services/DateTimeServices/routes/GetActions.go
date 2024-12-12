package routes

import (
	"date-time-service/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get Actions of Date Time services
// @Summary Get Actions from Date Time services
// @Description Get Actions from Date Time services
// @Tags Actions Date Time services
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string "Reactions name with parameters of it as object"
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
