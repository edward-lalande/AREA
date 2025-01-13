package routes

import (
	area "google/Area"
	_ "google/docs"
	"google/oauth"
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

	r.GET("/actions", area.GetActions)

	r.POST("/action", area.StoreActions)

	r.GET("/reactions", area.GetReactions)

	r.POST("/trigger", area.Trigger)
	r.POST("/reaction", area.StoreReactions)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
