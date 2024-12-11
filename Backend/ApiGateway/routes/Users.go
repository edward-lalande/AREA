package routes

import (
	models "api-gateway/Models"
	"api-gateway/utils"
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get User services
// @Summary Get to the user services
// @Description Routes get user services
// @Tags User api-gateway
// @Accept json
// @Produce json
// @Param Object body models.UsersGet true "routes wanted to the user services"
// @Success 200 {object} map[string]string "User services responses"
// @Failure 400 {object} map[string]string "Bad requests"
// @Failure 500 {object} map[string]string "Internal error"
// @Router /user [get]
func UserGet(c *gin.Context) {
	var body models.UsersGet
	c.ShouldBindJSON(&body)
	value := utils.GetEnvKey("USER_API")

	resp, err := http.Get(value + body.RoutesWanted)
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

// Post User services
// @Summary Post a new users to the user database without OAUTH2 or login
// @Description Routes to add a new user to the database
// @Tags User api-gateway
// @Accept json
// @Produce json
// @Param Object body models.UserInformation true "user information to login or sign-up"
// @Success 200 {string} string "User Token"
// @Failure 400 {object} map[string]string "Bad requests"
// @Failure 500 {object} map[string]string "Internal error"
// @Router /user [post]
func UserPost(c *gin.Context) {
	var body models.UserInformation
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	value := utils.GetEnvKey("USER_API")

	jsonBody, err := json.Marshal(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp, err := http.Post(value+body.RoutesWanted, "application/json", bytes.NewBuffer(jsonBody))
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
