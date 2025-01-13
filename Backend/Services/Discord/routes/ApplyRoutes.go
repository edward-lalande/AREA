package routes

import (
	_ "discord-service/docs"
	"discord-service/oauth"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func ApplyRoutes(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"ping": "pong"})
	})

	r.GET("/oauth", oauth.OAuthFront)
	r.GET("/add-oauth", oauth.AddOAuthFront)
	r.GET("/callback", oauth.CallBack)
	r.GET("/add-callback", oauth.AddCallBack)
	r.POST("/access-token", oauth.GetAccessToken)
	r.POST("/add-access-token", oauth.AddAccessToken)
	r.POST("/register", RegisterToken)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/action-name", GetActionsName)
	r.GET("/reaction-name", GetReactionsName)

	r.GET("/reactions", GetReactions)
	r.GET("/actions", GetActions)

	r.POST("/reaction", ReceivedReactions)
	r.POST("/action", RegisterAction)

	r.POST("/trigger", Trigger)
}
