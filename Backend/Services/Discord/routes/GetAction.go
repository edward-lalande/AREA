package routes

import (
	"discord-service/utils"
	"net/http"

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
