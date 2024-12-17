package area

import (
	"net/http"
	models "spotify/Models"
	"spotify/utils"

	"github.com/gin-gonic/gin"
)

func Actions(c *gin.Context) {
	var receivedData models.ActionsData
	db := utils.OpenDB(c)

	if err := c.ShouldBindJSON(&receivedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if db == nil {
		c.JSON(http.StatusInternalServerError, "Unable to open the database")
		return
	}

	db.Exec(c, "INSERT INTO \"SpotifyActions\" (area_id, action_type, user_token, is_playing, music_name)"+
		" VALUES ($1, $2, $3, $4, $5)", receivedData.AreaId, receivedData.ActionType, receivedData.AccessToken, receivedData.IsPlaying, receivedData.MusicName)

	defer db.Close(c)
	c.JSON(http.StatusAccepted, "Spotify Actions Accepted")
}
