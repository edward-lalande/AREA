package routes

import (
	models "github/Models"
	utils "github/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetWebhooks(c *gin.Context) {

	var (
		data models.Webhook
	)

	db := utils.OpenDB(c)

	if db == nil {
		return
	}

	rows, err := db.Query(c, "SELECT * FROM \"GithubWebhook\"")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	webhooks := []models.Webhook{}

	for rows.Next() {
		if err := rows.Scan(&data.Repository.Name, &data.Sender.Login, &data.HeadCommit.Timestamp); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse reaction"})
			return
		}
		webhooks = append(webhooks, data)
	}

	c.JSON(http.StatusOK, webhooks)
}
