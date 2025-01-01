package routes

import (
	area "dropbox/Area"
	"dropbox/oauth"
	"dropbox/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"ping": "pong"})
	})

	r.GET("/actions", func(c *gin.Context) {
		c.JSON(http.StatusOK, nil)
	})

	r.POST("/reaction", area.StoreReactions)
	r.GET("/reactions", func(c *gin.Context) {
		b, err := utils.OpenFile("Models/Reactions.json")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		json := utils.BytesToJson(b)
		c.JSON(http.StatusOK, json)
	})

	r.POST("/trigger", area.Trigger)

	r.GET("/oauth", oauth.OAuthFront)
	r.GET("/callback", oauth.CallBack)
	r.POST("/access-token", oauth.GetAccessToken)
}
