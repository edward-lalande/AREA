package routes

import (
	"context"
	"net/http"
	models "poc-crud-users/Models"
	"poc-crud-users/utils"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func isUserExists(receivedData models.Login, db *pgx.Conn) bool {
	var count int

	row := db.QueryRow(context.Background(), "SELECT COUNT(*) FROM \"User\" WHERE mail = $1 AND password = $2",
		receivedData.Mail, receivedData.Password)
	if err := row.Scan(&count); err != nil {
		return false
	}

	return count == 1
}

func LoginUserHandler(c *gin.Context) {
	var receivedData models.Login
	var user models.User
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

	token, err := utils.CreateToken(receivedData.Mail)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token, "test: ": "okok"})
}
