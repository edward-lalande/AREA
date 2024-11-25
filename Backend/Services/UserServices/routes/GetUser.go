package routes

import (
	"net/http"
	"poc-crud-users/utils"

	"github.com/gin-gonic/gin"
)

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
