package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.Engine) {

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"ping": "pong"})
	})

	r.GET("/actions", func(c *gin.Context) {
		c.JSON(http.StatusOK, nil)
	})

	//r.GET("/actions", GetActions)

	r.GET("/reactions", GetReactions)

	r.POST("/reaction", Reaction)

	r.POST("/trigger", Trigger)
}
