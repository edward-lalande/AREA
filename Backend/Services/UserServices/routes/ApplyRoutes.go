package routes

import "github.com/gin-gonic/gin"

func ApplyRoutes(r *gin.Engine) {
	r.GET("/ping", Ping)
	r.GET("/users", GetUser)
	r.POST("/sign-up", SignUpUserHandler)
	r.POST("/update", UpdateUser)
	r.POST("/login", LoginUserHandler)
	r.DELETE("/user", DeleteUser)
}
