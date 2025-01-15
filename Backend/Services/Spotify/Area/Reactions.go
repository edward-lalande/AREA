package area

import (
	"net/http"
	models "spotify/Models"
	"spotify/utils"

	"github.com/gin-gonic/gin"
)

func pauseSound(information models.Reactions) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("PUT", utils.GetEnvKey("SPOTIFY_PLAYER_API")+"me/player/pause", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+information.UserToken)
	req.Header.Set("Content-Type", "application/json")
	return client.Do(req)
}

// func startResume(information models.Reactions) (*http.Response, error) {
// 	client := &http.Client{}
// 	req, err := http.NewRequest("PUT", utils.GetEnvKey("SPOTIFY_PLAYER_API")+"me/player/play", nil)
// 	if err != nil {
// 		return nil, err
// 	}
// 	req.Header.Set("Authorization", "Bearer "+information.UserToken)
// 	req.Header.Set("Content-Type", "application/json")

// 	return client.Do(req)
// }

// func shuffleMode(information models.Reactions) (*http.Response, error) {
// 	client := &http.Client{}
// 	req, err := http.NewRequest("PUT", utils.GetEnvKey("SPOTIFY_PLAYER_API")+"me/player/shuffle?state=true", nil)
// 	if err != nil {
// 		return nil, err
// 	}
// 	req.Header.Set("Authorization", "Bearer "+information.UserToken)
// 	req.Header.Set("Content-Type", "application/json")

// 	return client.Do(req)
// }

func FindReactions(id int, information models.Reactions) (*http.Response, error) {
	reactions := map[int]func(models.Reactions) (*http.Response, error){
		0: pauseSound,
	}

	return reactions[id](information)
}

// Spotify Services
// @Summary Register an received Reactions
// @Description Register the reactions received by the message brocker with all informations nedded
// @Tags Spotify Area
// @Accept json
// @Produce json
// @Param routes body models.ReactionsReceived true "It must contains the AreaId and the reactions type"
// @Success 200 {object} map[string]string "Response is the received data"
// @Failure 400 {object} map[string]string "Invalid request it contains the error"
// @Failure 500 {object} map[string]string "Internal error it contains the error"
// @Router /reaction [post]
func ReceivedReactions(c *gin.Context) {
	var receivedData models.ReactionsReceived

	db := utils.OpenDB(c)
	if err := c.ShouldBindJSON(&receivedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := db.Exec(c, "INSERT INTO \"SpotifyReactions\" (area_id, reaction_type, user_token)"+
		" VALUES($1, $2, $3)", receivedData.AreaId, receivedData.ReactionType, receivedData.UserToken)
	if err != nil {
		c.JSON(http.StatusInsufficientStorage, gin.H{"error": err.Error()})
		return
	}

	defer db.Close(c)
	c.JSON(http.StatusAccepted, gin.H{"Spotify received": receivedData})
}

// Spotify Reactions
// @Summary send all the reactions
// @Description send all the reactions available on the Spotify services as an object arrays with the names and the object needed
// @Tags Spotify Area
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string "Response is the received data"
// @Failure 400 {object} map[string]string "Invalid request it contains the error"
// @Failure 500 {object} map[string]string "Internal error it contains the error"
// @Router /reactions [get]
func GetReactions(c *gin.Context) {
	b, err := utils.OpenFile(models.ReactionsModelsPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	json := utils.BytesToJson(b)
	c.JSON(http.StatusOK, json)
}
