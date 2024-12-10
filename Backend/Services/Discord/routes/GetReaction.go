package routes

import (
	models "discord-service/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetReactions(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"reactions": gin.H{"Send message on channel":models.ReactionGet{}}})
}
