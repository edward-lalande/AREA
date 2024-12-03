package oauth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CallBack(c *gin.Context) {
	code, _ := c.GetQuery("code")
	c.JSON(http.StatusOK, gin.H{"code": code})
}
