package routes

import (
	"dropbox/oauth"
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

	r.GET("/reactions", func(c *gin.Context) {
		c.JSON(http.StatusOK, nil)
	})

	r.GET("/oauth", oauth.OAuthFront)
	r.GET("/callback", oauth.CallBack)
	r.POST("/access-token", oauth.GetAccessToken)
}
