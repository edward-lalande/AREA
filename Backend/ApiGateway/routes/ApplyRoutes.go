package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type PingFrom struct {
	UserServices     bool
	DateTimeServices bool
}

// @title Tag Service API
func ApplyRoutes(r *gin.Engine) {
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/ping", Ping)

	r.GET("/user", userGet)
	r.POST("/user", userPost)

	r.GET("/time", GetTime)
}
