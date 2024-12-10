package routes

import (
	models "date-time-service/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get Actions of Date Time services
// @Summary Get Actions from Date Time services
// @Description Get Actions from Date Time services
// @Tags Actions Date Time services
// @Accept json
// @Produce json
// @Success 200 {object} models.GetTimeAction "Reactions name with parameters of it as object"
// @Router /actions [get]
func GetActions(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"actions": gin.H{"At time": models.GetTimeAction{1, 0, "", "", 0, 0}}})
}
