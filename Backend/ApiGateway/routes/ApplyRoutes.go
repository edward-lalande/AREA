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
	r.DELETE("/areas", DeleteArea)
	r.GET("/actions", GetActions)
	r.GET("/reactions", GetReactions)
	r.POST("/gitlab-webhook", GitlabWebhook)

	r.GET("/about.json", About)

	r.POST("/login", UserLogin)
	r.POST("/sign-up", UserSignUp)
	r.POST("/update-user", UserUpdate)

	r.GET("/discord/oauth", DiscordOauth2)
	r.GET("/spotify/oauth", SpotifyOauth2)
	r.GET("/github/oauth", GithubOauth2)
	r.GET("/gitlab/oauth", GitlabOauth2)
	r.GET("/google/oauth", GoogleOauth2)
	r.GET("/dropbox/oauth", DropBoxOauth2)
	r.GET("/asana/oauth", AsanaOauth2)
	r.GET("/miro/oauth", MiroOauth2)

	r.POST("/discord/access-token", DiscordAccessToken)
	r.POST("/spotify/access-token", SpotifyAccessToken)
	r.POST("/github/access-token", GithubAccessToken)
	r.POST("/gitlab/access-token", GitlabAccessToken)
	r.POST("/google/access-token", GoogleAccessToken)
	r.POST("/dropbox/access-token", DropboxAccessToken)
	r.POST("/asana/access-token", AsanaAccessToken)
	r.POST("/miro/access-token", MiroAccessToken)

	r.POST("/github-webhook", GithubWebhook)
}
