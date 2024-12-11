package routes

import (
	_ "date-time-service/docs"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func ApplyRoutes(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"ping": "pong"})
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/actions", GetActions)

	r.GET("/reactions", func(c *gin.Context) {
		c.JSON(http.StatusAccepted, nil)
	})

	r.POST("/action", RegisterAction)
	r.POST("/time", PostTime)
}
