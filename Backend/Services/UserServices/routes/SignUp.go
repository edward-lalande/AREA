package routes

import (
	"context"
	"net/http"
	"poc-crud-users/utils"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type SignUp struct {
	Mail     string `json:"mail"`
	Password string `json:"password"`
	Name     string `json:"name"`
	LastName string `json:"lastname"`
}

func isUserAlreadyExists(receivedData SignUp, db *pgx.Conn) bool {
	var count int
	row := db.QueryRow(context.Background(), "SELECT COUNT(*) FROM \"User\" WHERE mail = $1", receivedData.Mail)
	if err := row.Scan(&count); err != nil {
		return false
	}
	return count > 0
}

func writeInDB(receivedData SignUp, db *pgx.Conn) error {
	_, dbExecError := db.Exec(context.Background(), "INSERT INTO \"User\"(mail, password, name, lastname)"+
		" VALUES ($1, $2, $3, $4)", receivedData.Mail, receivedData.Password, receivedData.Name, receivedData.LastName)
	return dbExecError
}

func SignUpUserHandler(c *gin.Context) {
	var receivedData SignUp
	var db *pgx.Conn = utils.OpenDB(c)
	if db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to open Database"})
		return
	}

	if err := c.ShouldBindJSON(&receivedData); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	if isUserAlreadyExists(receivedData, db) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
		return
	}

	if writeValue := writeInDB(receivedData, db); writeValue != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": writeValue})
		return
	}

	db.Close(c)
	c.JSON(http.StatusOK, gin.H{"token": "bene"})
}
