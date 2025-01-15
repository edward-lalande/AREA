package routes

import (
	"context"
	"net/http"
	"poc-crud-users/utils"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
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
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(receivedData.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, dbExecError := db.Exec(context.Background(), "INSERT INTO \"User\"(mail, password, name, lastname)"+
		" VALUES ($1, $2, $3, $4)", receivedData.Mail, string(hashedPassword), receivedData.Name, receivedData.LastName)
	return dbExecError
}

// SignUpUserHandler registers a new user and generates an authentication token.
//
// @Summary User registration
// @Description Registers a new user by saving their details in the database. If successful, a JWT token is generated and returned.
// @Tags Authentication
// @Accept json
// @Produce json
// @Param user body SignUp true "User details for registration"
// @Success 200 {object} map[string]interface{} "JWT token for authentication"
// @Failure 400 {object} map[string]interface{} "Bad Request: User already exists or invalid input"
// @Failure 502 {object} map[string]interface{} "Bad Gateway: Error parsing request data"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /signup [post]
func SignUpUserHandler(c *gin.Context) {
	var (
		receivedData SignUp
		db           *pgx.Conn = utils.OpenDB(c)
	)

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

	var id string

	row := db.QueryRow(context.Background(), "SELECT id FROM \"User\" WHERE mail = $1", receivedData.Mail)
	_ = row.Scan(&id)
	defer db.Close(c)

	token, err := utils.CreateToken(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}
