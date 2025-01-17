package routes

import (
	area "gitlab/Area"
	"gitlab/oauth"
	"net/http"

	_ "gitlab/docs"

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
	r.GET("/add-oauth", oauth.AddOAuthFront)
	r.GET("/callback", oauth.CallBack)
	r.GET("/add-callback", oauth.AddCallBack)
	r.POST("/access-token", oauth.GetAccessToken)
	r.POST("/add-access-token", oauth.AddAccessToken)

	r.POST("/trigger", Trigger)

	r.GET("/reactions", area.GetReactions)

	r.POST("/reaction", area.StoreReactions)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
