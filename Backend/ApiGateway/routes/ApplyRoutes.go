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

func ApplyRoutes(r *gin.Engine) {
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/ping", Ping)

	r.GET("/services", Services)

	r.POST("/area", Area)

	r.GET("/user", UserGet)
	r.POST("/user", UserPost)

	r.GET("/discord", DiscordGet)
	r.POST("/discord", DiscordPost)

	r.GET("/time", GetTime)
	r.POST("/time", PostTime)
}
