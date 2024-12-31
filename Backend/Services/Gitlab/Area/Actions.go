package area

import (
	models "gitlab/Models"
	"gitlab/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Actions(c *gin.Context) {
	db := utils.OpenDB(c)
	receivedData := models.Actions{}

	if err := c.ShouldBindJSON(&receivedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	query := `INSERT INTO "GitlabActions" (action_type, area_id) VALUES ($1, $2)`
	_, err := db.Exec(c, query, receivedData.ActionType, receivedData.AreaId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert data"})
		return
	}

	defer db.Close(c)
	c.JSON(http.StatusOK, gin.H{"message": "Data inserted successfully"})
}
