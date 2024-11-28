package routes

import (
	"fmt"
	models "message-brocker/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendActionToService(c *gin.Context) {
	var receivedData models.ReceivedAction

	if err := c.ShouldBindJSON(&receivedData); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(receivedData)
}
