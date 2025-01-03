package routes

import (
	"net/http"
	area "spotify/Area"
	models "spotify/Models"
	"spotify/utils"

	"github.com/gin-gonic/gin"
)

// Spotify Services
// @Summary Trigger an Area
// @Description Actions triggerd the reactions and call the trigger route
// @Tags Spotify trigger
// @Accept json
// @Produce json
// @Param routes body models.TriggerdModels true "It contains the Area Id to the reactions"
// @Success 200 {object} map[string]string "Response of the reactions"
// @Failure 400 {object} map[string]string "Invalid request it contains the error"
// @Failure 500 {object} map[string]string "Internal error it contains the error"
// @Router /trigger [post]
func Trigger(c *gin.Context) {
	var (
		receivedData models.TriggerdModels
		user         models.Reactions
	)

	if err := c.ShouldBindJSON(&receivedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := utils.OpenDB(c)
	row := db.QueryRow(c, "SELECT reaction_type FROM \"SpotifyReactions\" WHERE area_id = $1", receivedData.AreaId)

	if err := row.Scan(&user.ReactionType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer db.Close(c)

	rep, _ := area.FindReactions(user.ReactionType, user)
	c.JSON(rep.StatusCode, gin.H{
		"body": rep.Body,
	})
}
