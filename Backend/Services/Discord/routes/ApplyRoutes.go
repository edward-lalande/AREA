package routes

import (
	"discord-service/oauth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"ping": "pong"})
	})

	r.GET("/oauth2", oauth.OAuthFront)
	r.GET("/callback", oauth.CallBack)
	r.POST("/access-token", oauth.GetAccessToken)
	r.POST("/register", RegisterToken)

	r.POST("/reaction", ReceivedReactions)
	r.POST("/active-reactions", ActiveReactions)
	r.POST("/trigger", Trigger)
}
