package routes

import (
	models "discord-service/Models"
	"discord-service/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Discord Reactions
// @Summary send all the reactions
// @Description send all the reactions available on the discord services as an object arrays with the names and the object needed
// @Tags Discord Area
// @Accept json
// @Produce json
// @Success 200 {object} models.ReactionGet "Response is the received data"
// @Failure 400 {object} map[string]string "Invalid request it contains the error"
// @Failure 500 {object} map[string]string "Internal error it contains the error"
// @Router /reactions [get]
func GetReactions(c *gin.Context) {
	b, err := utils.OpenFile(models.ReactionModelPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	json := utils.BytesToJson(b)

	c.JSON(http.StatusOK, json)
}
