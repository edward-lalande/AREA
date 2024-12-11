package routes

import (
	"net/http"
	"poc-crud-users/utils"

	"github.com/gin-gonic/gin"
)

// GetUser retrieves a user's details from the database.
//
// @Summary Retrieve user information
// @Description Fetches the details of a user from the "User" table in the database.
// @Tags Users
// @Produce json
// @Success 200 {object} []SignUp "Successfully retrieved user information"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /users [get]
func GetUser(c *gin.Context) {
	var value []SignUp
	db := utils.OpenDB(c)
	if db == nil {
		return
	}

	row := db.QueryRow(c, "SELECT * FROM \"User\"")
	err := row.Scan(value)
	if err == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err})
	}
	db.Close(c)
	c.JSON(http.StatusOK, value)
}
