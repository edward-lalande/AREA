package oauth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Gitlab OAUTH2
// @Summary Send the code received by Gitlab to the frontend
// @Description Send the code received by Gitlab to the frontend
// @Tags Gitlab OAUTH2
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string "the code to redirect to"
// @Router /callback [get]
func CallBack(c *gin.Context) {
	code, exists := c.GetQuery("code")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing code"})
		return
	}
	redirectURL := "http://localhost:8081/login?gitlab_code=" + code
	c.Redirect(http.StatusFound, redirectURL)
}

func AddCallBack(c *gin.Context) {
	code, exists := c.GetQuery("code")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing code"})
		return
	}
	redirectURL := "http://localhost:8081/account?gitlab_code=" + code
	c.Redirect(http.StatusFound, redirectURL)
}
