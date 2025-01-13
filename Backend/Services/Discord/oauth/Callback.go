package oauth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Discord OAUTH2
// @Summary Send the code received by discord to the frontend
// @Description Send the code received by discord to the frontend
// @Tags Discord OAUTH2
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string "the code to redirect to"
// @Router /callback [get]
func CallBack(c *gin.Context) {
	token := c.GetHeader("token")

	fmt.Println("discord token => [" + token + "]")
	code, exists := c.GetQuery("code")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing code"})
		return
	}
	redirectURL := "http://localhost:8081/login?discord_code=" + code
	c.Redirect(http.StatusFound, redirectURL)
}

func AddCallBack(c *gin.Context) {
	token := c.GetHeader("token")

	fmt.Println("token => [" + token + "]")
	code, exists := c.GetQuery("code")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing code"})
		return
	}
	redirectURL := "http://localhost:8081/account?discord_code=" + code
	c.Redirect(http.StatusFound, redirectURL)
}
