package routes

import (
	models "discord-service/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)
// Discord Reactions
// @Summary send all the reactions
// @Description send all the reactions available on the discord services as an object arrays with the names and the object needed
// @Tags Discord Area
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string "Response is the received data"
// @Failure 400 {object} map[string]string "Invalid request it contains the error"
// @Failure 500 {object} map[string]string "Internal error it contains the error"
// @Router /reactions [get]
func GetReactions(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"reactions": gin.H{"Send message on channel":models.ReactionGet{}}})
}
