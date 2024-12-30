package routes

import (
	_ "api-gateway/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func ApplyRoutes(r *gin.Engine) {
	r.GET("/ping", Ping)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/services", Services)

	r.POST("/areas", Area)
	r.GET("/areas", GetUserAreas)
	r.GET("/actions", GetActions)
	r.GET("/reactions", GetReactions)

	r.POST("/login", UserLogin)
	r.POST("/sign-up", UserSignUp)
	r.POST("/update-user", UserUpdate)

	r.GET("/discord/oauth", DiscordOauth2)
	r.GET("/spotify/oauth", SpotifyOauth2)
	r.GET("/github/oauth", GithubOauth2)
	r.GET("/gitlab/oauth", GitlabOauth2)

	r.POST("/discord/access-token", DiscordAccessToken)
	r.POST("/spotify/access-token", SpotifyAccessToken)
	r.POST("/github/access-token", GithubAccessToken)
	r.POST("/gitlab/access-token", GitlabAccessToken)

	r.POST("/github-webhook", GithubWebhook)
}
