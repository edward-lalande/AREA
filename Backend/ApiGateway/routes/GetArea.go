package routes

import (
	"api-gateway/utils"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type actionsList struct {
	name string
	url  string
}

type reactionsList struct {
	name string
	url  string
}

// Get Actions of all services
// @Summary Get actions from all services
// @Description Get actions from all services
// @Tags Area api-gateway
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string "Actions name with parameters of it as object"
// @Failure 500 {object} map[string]string "Internal error"
// @Router /actions [get]
func GetActions(c *gin.Context) {
	var actions []map[string]interface{}

	servicesArray := []actionsList{
		{"Date Time Services", utils.GetEnvKey("TIME_API")},
		{"Discord Services", utils.GetEnvKey("DISCORD_API")},
		{"Gitlab Services", utils.GetEnvKey("GITLAB_API")},
		{"Spotify Services", utils.GetEnvKey("SPOTIFY_API")},
		{"Google Services", utils.GetEnvKey("GOOGLE_API")},
		{"Github Services", utils.GetEnvKey("GITHUB_API")},
		{"Meteo Services", utils.GetEnvKey("METEO_API")},
		{"Ticket Master Services", utils.GetEnvKey("TICKET_MASTER_API")},
		{"Cryptomoney Services", utils.GetEnvKey("CRYPTOMONEY_API")},
	}

	for _, service := range servicesArray {
		resp, err := http.Get(service.url + "actions")
		if err != nil {
			continue
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			continue
		}

		var jsonResponse map[string]interface{}
		if err := json.Unmarshal(body, &jsonResponse); err != nil {
			continue
		}
		actions = append(actions, jsonResponse)
	}

	c.JSON(http.StatusOK, actions)
}

// Get Reactions of all services
// @Summary Get reactions from all services
// @Description Get reactions from all services
// @Tags Area api-gateway
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string "Reactions name with parameters of it as object"
// @Failure 500 {object} map[string]string "Internal error"
// @Router /reactions [get]
func GetReactions(c *gin.Context) {
	var reactions []map[string]interface{}

	servicesArray := []reactionsList{
		{"Discord Services", utils.GetEnvKey("DISCORD_API")},
		{"Spotify Services", utils.GetEnvKey("SPOTIFY_API")},
		{"Gitlab Services", utils.GetEnvKey("GITLAB_API")},
		{"Google Services", utils.GetEnvKey("GOOGLE_API")},
		{"Github Services", utils.GetEnvKey("GITHUB_API")},
		{"DropBox Services", utils.GetEnvKey("DROPBOX_API")},
		{"Asana Services", utils.GetEnvKey("ASANA_API")},
		{"Miro Services", utils.GetEnvKey("Miro_API")},
	}

	for _, service := range servicesArray {
		resp, err := http.Get(service.url + "reactions")
		if err != nil {
			continue
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			continue
		}

		var jsonResponse map[string]interface{}
		if err := json.Unmarshal(body, &jsonResponse); err != nil {
			continue
		}

		reactions = append(reactions, jsonResponse)
	}

	c.JSON(http.StatusOK, reactions)
}
