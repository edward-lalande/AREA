package routes

import (
	_ "github/docs"
	"github/oauth"
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

	r.POST("/action", createAction)

	r.GET("/actions", getActions)

	r.GET("/reactions", func(c *gin.Context) {
		c.JSON(http.StatusOK, nil)
	})

	r.POST("/webhook/push", GetWebhooksPush)
	r.POST("/webhook/commit_comment", GetWebhooksCommitComment)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
