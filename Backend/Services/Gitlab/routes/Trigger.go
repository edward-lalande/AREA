package routes

import (
	area "gitlab/Area"
	models "gitlab/Models"
	"gitlab/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Trigger(c *gin.Context) {
	var (
		receivedData models.TriggerdModels
		database     models.Database
	)

	if err := c.ShouldBindJSON(&receivedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := utils.OpenDB(c)
	row := db.QueryRow(c, "SELECT user_token, reaction_type, project_id, body FROM \"GitlabReactions\" WHERE area_id = $1", receivedData.AreaId)

	if err := row.Scan(&database.UserToken, &database.ReactionType, &database.ProjectId, &database.Body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer db.Close(c)

	rep, err := area.FindReaction(c, database)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(rep.StatusCode, gin.H{
		"body": rep.Body,
	})
}
