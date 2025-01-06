package oauth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Miro OAUTH2
// @Summary Send the code received by Miro to the frontend
// @Description Send the code received by Miro to the frontend
// @Tags Miro OAUTH2
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
	redirectURL := "http://localhost:8081/login?miro_code=" + code
	c.Redirect(http.StatusFound, redirectURL)
}
