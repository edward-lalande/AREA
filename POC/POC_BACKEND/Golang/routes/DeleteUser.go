package routes

import (
	"context"
	"net/http"
	"poc-crud-users/utils"

	"github.com/gin-gonic/gin"
)

type LoginId struct {
	Id int `json:"id"`
}

func DeleteUser(c *gin.Context) {
	var receivedData LoginId
	var user User
	db := utils.OpenDB(c)

	if db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to open Database"})
		return
	}
	if err := c.ShouldBindJSON(&receivedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	row := db.QueryRow(context.Background(), "DELETE FROM \"User\" WHERE id = $1",
		receivedData.Id)
	_ = row.Scan(&user.Id)

	db.Close(c)
	c.JSON(http.StatusOK, gin.H{"user_deleted": receivedData.Id})
}
