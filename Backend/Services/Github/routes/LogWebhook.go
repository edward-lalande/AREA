package routes

import (
	models "github/Models"
	utils "github/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateWebhook(c *gin.Context) {

	var (
		data models.Webhook
	)

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := utils.OpenDB(c)

	if db == nil {
		return
	}

	query := `INSERT INTO "GithubWebhook" (name, login, timestamp) VALUES ($1, $2, $3) RETURNING id;`

	_, err := db.Exec(c, query, data.Repository.Name, data.Sender.Login, data.HeadCommit.Timestamp)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	defer db.Close(c)
}
