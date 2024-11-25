package routes

import (
	"context"
	"net/http"
	models "poc-crud-users/Models"
	"poc-crud-users/utils"

	"github.com/gin-gonic/gin"
)

func UpdateUser(c *gin.Context) {
	var receivedData models.UserInformation
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

	usefullInformation := models.Login{receivedData.Mail, receivedData.Password}
	if !isUserExists(usefullInformation, db) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Wrong email or password"})
		return
	}

	row := db.QueryRow(context.Background(), "ALTER \"User\" $1 WITH password $2",
		receivedData.Id, receivedData.Password)
	_ = row.Scan(&user.Id)
	row = db.QueryRow(context.Background(), "ALTER \"User\" $1 WITH mail $2",
		receivedData.Id, receivedData.Mail)
	_ = row.Scan(&user.Id)
	row = db.QueryRow(context.Background(), "ALTER \"User\" $1 WITH name $2",
		receivedData.Id, receivedData.Name)
	_ = row.Scan(&user.Id)
	row = db.QueryRow(context.Background(), "ALTER \"User\" $1 WITH lastname $2",
		receivedData.Id, receivedData.Lastname)
	_ = row.Scan(&user.Id)

	db.Close(c)
	c.JSON(http.StatusOK, gin.H{"user_updated": receivedData.Id})
}
