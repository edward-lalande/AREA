package routes

import (
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
	r.GET("/callback", oauth.CallBack)
	r.POST("/access-token", oauth.GetAccessToken)

	r.POST("/action", createAction)

	r.GET("/actions", getActions)

	r.GET("/reactions", func(c *gin.Context) {
		c.JSON(http.StatusOK, nil)
	})

	r.POST("/webhook/push", GetWebhooksPush)
	r.POST("/webhook/commit_comment", GetWebhooksCommitComment)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
