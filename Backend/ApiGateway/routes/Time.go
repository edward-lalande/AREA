package routes

import (
	models "api-gateway/Models"
	"api-gateway/utils"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTime(c *gin.Context) {
	var body models.DateTimeResponse
	c.ShouldBindJSON(&body)
	fmt.Println(utils.GetEnvKey("TIME_API") + body.Routes)
	resp, err := http.Get(utils.GetEnvKey("TIME_API") + body.Routes)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(resp.StatusCode, gin.H{
		"headers": resp.Header,
		"body":    utils.BytesToJson(respBody),
	})
}
