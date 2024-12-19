package area

import (
	"fmt"
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
	fmt.Println("id: ", id)
	reactions := map[int]func(models.Reactions) (*http.Response, error){
		0: pauseSound,
	}

	return reactions[id](information)
}

func ReceivedReactions(c *gin.Context) {
	var receivedData models.ReactionsReceived

	db := utils.OpenDB(c)
	if err := c.ShouldBindJSON(&receivedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("received type: ", receivedData.ReactionType)
	_, err := db.Exec(c, "INSERT INTO \"SpotifyReactions\" (area_id, reaction_type, user_token)"+
		" VALUES($1, $2, $3)", receivedData.AreaId, receivedData.ReactionType, receivedData.UserToken)
	if err != nil {
		c.JSON(http.StatusInsufficientStorage, gin.H{"error": err.Error()})
		return
	}

	defer db.Close(c)
	c.JSON(http.StatusAccepted, gin.H{"Discord received": receivedData})
}
