package routes

import (
	"context"
	"net/http"
	"poc-crud-users/utils"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type Login struct {
	Mail     string `json:"mail"`
	Password string `json:"password"`
}

type User struct {
	Id    int    `json:"id"`
	Login string `json:"name"`
}

func isUserExists(receivedData Login, db *pgx.Conn) bool {
	var count int

	row := db.QueryRow(context.Background(), "SELECT COUNT(*) FROM \"User\" WHERE mail = $1 AND password = $2",
		receivedData.Mail, receivedData.Password)
	if err := row.Scan(&count); err != nil {
		return false
	}

	return count == 1
}

func LoginUserHandler(c *gin.Context) {
	var receivedData Login
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

	if !isUserExists(receivedData, db) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Wrong email or password"})
		return
	}

	row := db.QueryRow(context.Background(), "SELECT id, name FROM \"User\" WHERE mail = $1 AND password = $2",
		receivedData.Mail, receivedData.Password)
	_ = row.Scan(&user.Id, &user.Login)

	db.Close(c)
	c.JSON(http.StatusOK, user)
}
