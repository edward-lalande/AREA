package routes

import (
	"fmt"
	"net/http"
	models "poc-crud-users/Models"
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
	var user models.User

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

	fmt.Println(id)

	row := db.QueryRow(c, "SELECT * FROM \"User\" WHERE id = $1", id)
	err := row.Scan(&user.Id, &user.Mail, &user.Password, &user.Login, &user.Lastname, &user.AsanaToken, &user.DiscordToken,
		&user.DropboxToken, &user.GithubToken, &user.GitlabToken, &user.GoogleToken, &user.MiroToken, &user.SpotifyToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}
