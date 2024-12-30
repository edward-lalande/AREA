package routes

import (
	"meteo/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"ping": "pong"})
	})

	r.GET("/actions", func(c *gin.Context) {
		b, err := utils.OpenFile("Models/Actions.json")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		json := utils.BytesToJson(b)
		c.JSON(http.StatusOK, json)
	})
	r.POST("/action", StoreActions)

}
