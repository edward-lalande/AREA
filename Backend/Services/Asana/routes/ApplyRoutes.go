package routes

import (
	area "asana/Area"
	_ "asana/docs"
	"asana/oauth"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func ApplyRoutes(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"ping": "pong"})
	})

	r.GET("/actions", func(c *gin.Context) {
		c.JSON(http.StatusOK, nil)
	})

	r.GET("/oauth", oauth.OAuthFront)
	r.GET("/callback", oauth.CallBack)
	r.POST("/access-token", oauth.GetAccessToken)

	r.GET("/reactions", area.GetReactions)

	r.POST("/reaction", area.StoreReactions)
	r.POST("/trigger", area.Trigger)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
