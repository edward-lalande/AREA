package routes

import (
	_ "api-gateway/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func ApplyRoutes(r *gin.Engine) {
	r.GET("/ping", Ping)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/services", Services)

	r.POST("/area", Area)
	r.GET("/areas", GetUserAreas)
	r.GET("/actions", GetActions)
	r.GET("/reactions", GetReactions)

	r.POST("/login", UserLogin)
	r.POST("/sign-up", UserSignUp)
	r.POST("/update-user", UserUpdate)
}
