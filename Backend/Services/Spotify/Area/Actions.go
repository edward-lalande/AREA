package area

import (
	"fmt"
	"io"
	"net/http"
	models "spotify/Models"
	"spotify/utils"

	"github.com/gin-gonic/gin"
)

func getUserSpotifyId(spotifyToken string) string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", utils.GetEnvKey("SPOTIFY_PLAYER_API")+"me", nil)
	if err != nil {
		return ""
	}
	req.Header.Set("Authorization", "Bearer "+spotifyToken)
	req.Header.Set("Content-Type", "application/json")
	resp, _ := client.Do(req)
	b, _ := io.ReadAll(resp.Body)
	json := utils.BytesToJson(b)
	defer resp.Body.Close()
	return json["id"].(string)
}

func GetNbPlaylists(spotifyToken, id string) int {
	client := &http.Client{}

	req, err := http.NewRequest("GET", utils.GetEnvKey("SPOTIFY_PLAYER_API")+"users/"+id+"/playlists", nil)
	if err != nil {
		return 0
	}

	req.Header.Set("Authorization", "Bearer "+spotifyToken)
	req.Header.Set("Content-Type", "application/json")
	resp, _ := client.Do(req)
	b, _ := io.ReadAll(resp.Body)
	json := utils.BytesToJson(b)
	defer resp.Body.Close()
	if json["total"] == nil {
		fmt.Println("playlists -1")
		return -1
	}
	fmt.Println("json total: ", json["total"])
	return int(json["total"].(float64))
}

// Post Spotify Actions
// @Summary Post an Actions
// @Description Post an Spotify actions, receive by the Message Brocker (handler of communication between services) and register it to him database
// @Tags Actions Spotify services
// @Accept json
// @Produce json
// @Param routes body models.ActionsData true "It contains the Area Id, the location and the Spotify of the Area"
// @Success 200 {object} map[string]string "Response is the Id of the Area"
// @Failure 400 {object} map[string]string "Invalid request"
// @Failure 500 {object} map[string]string "Internal error"
// @Router /actions [post]
func Actions(c *gin.Context) {
	var receivedData models.ActionsData
	var user models.User
	db := utils.OpenDB(c)

	if err := c.ShouldBindJSON(&receivedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if db == nil {
		c.JSON(http.StatusInternalServerError, "Unable to open the database")
		return
	}

	userId := utils.ParseToken(receivedData.AccessToken)
	if userId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token"})
		return
	}

	row := db.QueryRow(c, "SELECT * FROM \"User\" WHERE id = $1", userId)
	err := row.Scan(&user.Id, &user.Mail, &user.Password, &user.Login, &user.Lastname, &user.AsanaToken, &user.DiscordToken,
		&user.DropboxToken, &user.GithubToken, &user.GitlabToken, &user.GoogleToken, &user.MiroToken, &user.SpotifyToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	id := getUserSpotifyId(*user.SpotifyToken)
	nbPlaylists := GetNbPlaylists(receivedData.AccessToken, id)

	db.Exec(c, "INSERT INTO \"SpotifyActions\" (area_id, action_type, user_token, user_id, is_playing, nb_playlists)"+
		" VALUES ($1, $2, $3, $4, $5, $6)", receivedData.AreaId, receivedData.ActionType, *user.SpotifyToken, id, receivedData.IsPlaying, nbPlaylists)

	defer db.Close(c)
	c.JSON(http.StatusAccepted, "Spotify Actions Accepted")
}

// Get Actions of Spotify
// @Summary Get Actions from Spotify
// @Description Get Actions from Spotify
// @Tags Actions Spotify
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string "Reactions name with parameters of it as object"
// @Router /actions [get]
func GetActions(c *gin.Context) {
	b, err := utils.OpenFile(models.ActionsModelsPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	json := utils.BytesToJson(b)
	c.JSON(http.StatusOK, json)
}
