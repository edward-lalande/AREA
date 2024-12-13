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

// Post User services
// @Summary Update a user
// @Description Update a user to the user services database
// @Tags User api-gateway
// @Accept json
// @Produce json
// @Param Object body models.SignUp true "user information to update"
// @Success 200 {string} string "User Token"
// @Failure 400 {object} map[string]string "Bad requests"
// @Failure 500 {object} map[string]string "Internal error"
// @Router /update-user [post]
func UserUpdate(c *gin.Context) {
	var data models.SignUp
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	value := utils.GetEnvKey("USER_API")

	jsonBody, err := json.Marshal(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp, err := http.Post(value+"update", "application/json", bytes.NewBuffer(jsonBody))
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
// @Summary Login a user
// @Description Login a user to the user services database
// @Tags User api-gateway
// @Accept json
// @Produce json
// @Param Object body models.Login true "user information to login"
// @Success 200 {string} string "User Token"
// @Failure 400 {object} map[string]string "Bad requests"
// @Failure 500 {object} map[string]string "Internal error"
// @Router /login [post]
func UserLogin(c *gin.Context) {
	var data models.Login

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	value := utils.GetEnvKey("USER_API")

	jsonBody, err := json.Marshal(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp, err := http.Post(value+"login", "application/json", bytes.NewBuffer(jsonBody))
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
// @Summary Sign up a user
// @Description Sign up a user to the user services database
// @Tags User api-gateway
// @Accept json
// @Produce json
// @Param Object body models.SignUp true "user information to sign-up"
// @Success 200 {string} string "User Token"
// @Failure 400 {object} map[string]string "Bad requests"
// @Failure 500 {object} map[string]string "Internal error"
// @Router /sign-up [post]
func UserSignUp(c *gin.Context) {
	var data models.SignUp

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	value := utils.GetEnvKey("USER_API")

	jsonBody, err := json.Marshal(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp, err := http.Post(value+"sign-up", "application/json", bytes.NewBuffer(jsonBody))
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
