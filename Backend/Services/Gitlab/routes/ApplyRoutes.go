package routes

import (
	area "gitlab/Area"
	"gitlab/oauth"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func ApplyRoutes(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"ping": "pong"})
	})

	r.POST("/action", area.Actions)

	r.GET("/actions", area.GetActions)

	r.POST("/webhook", Webhook)

	r.GET("/oauth", oauth.OAuthFront)
	r.GET("/callback", oauth.CallBack)
	r.POST("/access-token", oauth.GetAccessToken)

	r.POST("/trigger", Trigger)

	r.GET("/reactions", area.GetReactions)

	r.POST("/reaction", area.StoreReactions)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
