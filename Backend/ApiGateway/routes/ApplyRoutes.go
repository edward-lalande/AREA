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
	r.GET("/ping", Ping)

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	r.GET("/services", Services)

	r.POST("/area", Area)

	r.POST("/webhooks-discord", DiscordWebHooks)
	r.GET("/user", UserGet)
	r.POST("/user", UserPost)

	r.GET("/discord", DiscordGet)
	r.POST("/discord", DiscordPost)

	r.GET("/time", GetTime)
}
