package routes

import "github.com/gin-gonic/gin"

func ApplyRoutes(r *gin.Engine) {
	r.GET("/ping", Ping)
	r.GET("/datetime", GetTime)
	r.GET("/users", GetUser)
	r.POST("/user", SignUpUserHandler)
	r.POST("/user-update", UpdateUser)
	r.POST("/login", LoginUserHandler)
	r.DELETE("/user", DeleteUser)
}
