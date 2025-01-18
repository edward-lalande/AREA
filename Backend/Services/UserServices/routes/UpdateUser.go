package routes

import (
	"net/http"
	models "poc-crud-users/Models"
	"poc-crud-users/utils"

	"github.com/gin-gonic/gin"
)

// UpdateUser updates an existing user's information in the database.
//
// @Summary Update user information
// @Description Updates a user's details such as email, password, first name, and last name in the "User" table.
// @Tags Users
// @Accept json
// @Produce json
// @Param user body models.UserInformation true "User information to update"
// @Success 200 {object} map[string]interface{} "Confirmation of the updated user"
// @Failure 400 {object} map[string]interface{} "Bad Request: Invalid input or user not found"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /user [post]
func UpdateName(c *gin.Context) {
	var receivedData models.UserInformation

	if err := c.ShouldBindJSON(&receivedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token := c.GetHeader("token")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token is required"})
		return
	}

	id := utils.ParseToken(token)
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token"})
		return
	}

	db := utils.OpenDB(c)
	if db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to open the database"})
		return
	}
	defer db.Close(c)

	_, err := db.Query(c, "UPDATE \"User\" SET name = $1 WHERE id = $2", receivedData.Name, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"info": "User updated!"})
}

func UpdateEmail(c *gin.Context) {
	var receivedData models.UserInformation

	if err := c.ShouldBindJSON(&receivedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token := c.GetHeader("token")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token is required"})
		return
	}

	id := utils.ParseToken(token)
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token"})
		return
	}

	db := utils.OpenDB(c)
	if db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to open the database"})
		return
	}
	defer db.Close(c)

	_, err := db.Query(c, "UPDATE \"User\" SET mail = $1 WHERE id = $2", receivedData.Mail, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"info": "User updated!"})
}

func UpdateLastname(c *gin.Context) {
	var receivedData models.UserInformation

	if err := c.ShouldBindJSON(&receivedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token := c.GetHeader("token")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token is required"})
		return
	}

	id := utils.ParseToken(token)
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token"})
		return
	}

	db := utils.OpenDB(c)
	if db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to open the database"})
		return
	}
	defer db.Close(c)

	_, err := db.Query(c, "UPDATE \"User\" SET lastname = $1 WHERE id = $2", receivedData.Lastname, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"info": "User updated!"})
}
