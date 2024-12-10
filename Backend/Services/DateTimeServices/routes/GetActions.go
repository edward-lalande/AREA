package routes

import (
	models "date-time-service/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetActions(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"actions": gin.H{"At time": models.GetTimeAction{}}})
}
