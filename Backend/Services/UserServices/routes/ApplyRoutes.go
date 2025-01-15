package routes

import (
	_ "poc-crud-users/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func ApplyRoutes(r *gin.Engine) {
	r.GET("/ping", Ping)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/user", GetUser)

	r.POST("/sign-up", SignUpUserHandler)
	r.POST("/login", LoginUserHandler)

	r.POST("/update-email", UpdateEmail)
	r.POST("/update-name", UpdateName)
	r.POST("/update-lastname", UpdateLastname)
	r.DELETE("/user", DeleteUser)
}
