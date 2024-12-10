package routes

import (
	"context"
	"net/http"
	models "poc-crud-users/Models"
	"poc-crud-users/utils"

	"github.com/gin-gonic/gin"
)

// DeleteUser deletes a user from the database.
//
// @Summary Delete a user
// @Description Deletes a user from the "User" table in the database using the provided user ID.
// @Tags Users
// @Accept json
// @Produce json
// @Param id body models.LoginId true "User ID"
// @Success 200 {object} map[string]interface{} "Confirmation of the deleted user"
// @Failure 400 {object} map[string]interface{} "Bad Request: Invalid input"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /user [delete]
func DeleteUser(c *gin.Context) {
	var receivedData models.LoginId
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

	row := db.QueryRow(context.Background(), "DELETE FROM \"User\" WHERE id = $1",
		receivedData.Id)
	_ = row.Scan(&user.Id)

	db.Close(c)
	c.JSON(http.StatusOK, gin.H{"user_deleted": receivedData.Id})
}
