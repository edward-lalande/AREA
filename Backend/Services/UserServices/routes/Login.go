package routes

import (
	"context"
	"net/http"
	models "poc-crud-users/Models"
	"poc-crud-users/utils"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

func isUserExists(receivedData models.Login, db *pgx.Conn) bool {
	var hashedPassword string
	row := db.QueryRow(context.Background(), "SELECT password FROM \"User\" WHERE mail = $1", receivedData.Mail)
	if err := row.Scan(&hashedPassword); err != nil {
		return false
	}

	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(receivedData.Password)) == nil
}

// LoginUserHandler logs in a user and generates an authentication token.
//
// @Summary User login
// @Description Authenticates a user by verifying their email and password, then generates a JWT token upon successful login.
// @Tags Authentication
// @Accept json
// @Produce json
// @Param credentials body models.Login true "User credentials"
// @Success 200 {object} map[string]interface{} "JWT token for authentication"
// @Failure 400 {object} map[string]interface{} "Bad Request: Invalid credentials"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /login [post]
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

	token, err := utils.CreateToken(user.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
