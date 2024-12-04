package routes

import (
	"context"
	models "discord-service/Models"
	"discord-service/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ActiveReactions(c *gin.Context) {
	var (
		body  models.ActiveReactionData
		count int
	)
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := utils.OpenDB(c)
	if db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to open database"})
		return
	}
	row := db.QueryRow(context.Background(), "SELECT * FROM \"Reactions\" WHERE service_id = $1 AND action_id = $2", body.ServiceId, body.ActionId)
	if err := row.Scan(&count); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}
