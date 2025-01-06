package routes

import (
	_ "api-gateway/docs"
	"api-gateway/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

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
	r.POST("/gitlab-webhook", func(c *gin.Context) {
		var a map[string]interface{}
		if err := c.ShouldBindJSON(&a); err != nil {
			fmt.Println("Error parsing JSON:", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
			return
		}

		jsonData, err := json.Marshal(a)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal error"})
			return
		}

		resp, err := http.Post(utils.GetEnvKey("GITLAB_API")+"webhook", "application/json", bytes.NewReader(jsonData))
		if err != nil {
			c.JSON(http.StatusInternalServerError, resp.Body)
			return
		}
		defer resp.Body.Close()

		c.JSON(http.StatusOK, gin.H{"message": "Webhook forwarded successfully"})
	})

	r.POST("/login", UserLogin)
	r.POST("/sign-up", UserSignUp)
	r.POST("/update-user", UserUpdate)

	r.GET("/discord/oauth", DiscordOauth2)
	r.GET("/spotify/oauth", SpotifyOauth2)
	r.GET("/github/oauth", GithubOauth2)
	r.GET("/gitlab/oauth", GitlabOauth2)
	r.GET("/google/oauth", GoogleOauth2)
	r.GET("/dropbox/oauth", DropBoxOauth2)

	r.POST("/discord/access-token", DiscordAccessToken)
	r.POST("/spotify/access-token", SpotifyAccessToken)
	r.POST("/github/access-token", GithubAccessToken)
	r.POST("/gitlab/access-token", GitlabAccessToken)
	r.POST("/google/access-token", GoogleAccessToken)
	r.POST("/dropbox/access-token", DropboxAccessToken)

	r.POST("/github-webhook", GithubWebhook)
}
