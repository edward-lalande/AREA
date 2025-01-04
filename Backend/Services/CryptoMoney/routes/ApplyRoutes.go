package routes

import (
	area "cryptomoney/Area"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"ping": "pong"})
	})

	r.POST("/action", area.StoreActions)
	r.GET("/actions", area.GetActions)
}
