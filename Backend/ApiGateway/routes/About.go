package routes

import (
	"api-gateway/utils"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Action struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Reaction struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ActionsResponse struct {
	Actions []Action `json:"actions"`
}

type ReactionsResponse struct {
	Reactions []Reaction `json:"reactions"`
}

type ServiceResponse struct {
	Name      string      `json:"name"`
	Actions   interface{} `json:"actions"`
	Reactions interface{} `json:"reactions"`
}

type Service struct {
	Name string
	Url  string
}

type IP struct {
	Query string
}

func getIp() string {
	req, err := http.Get("http://ip-api.com/json/")
	if err != nil {
		return err.Error()
	}
	defer req.Body.Close()

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return err.Error()
	}

	var ip IP
	json.Unmarshal(body, &ip)

	return ip.Query
}

func About(c *gin.Context) {
	servicesArray := []Service{
		{"Date Time", utils.GetEnvKey("TIME_API")},
		{"Discord", utils.GetEnvKey("DISCORD_API")},
		{"Dropbox", utils.GetEnvKey("DROPBOX_API")},
		{"Github", utils.GetEnvKey("GITHUB_API")},
		{"Gitlab", utils.GetEnvKey("GITLAB_API")},
		{"Google", utils.GetEnvKey("GOOGLE_API")},
		{"Meteo", utils.GetEnvKey("METEO_API")},
		{"Spotify", utils.GetEnvKey("SPOTIFY_API")},
		{"Asana", utils.GetEnvKey("ASANA_API")},
		{"Ticket Master", utils.GetEnvKey("TICKET_MASTER_API")},
		{"Twilio", utils.GetEnvKey("TWILIO_API")},
		{"Cryptomoney", utils.GetEnvKey("CRYPTOMONEY_API")},
		{"Miro", utils.GetEnvKey("MIRO_API")},
	}

	var servicesJson []ServiceResponse
	var actionsResponse interface{} = nil
	var reactionsResponse interface{} = nil

	for _, service := range servicesArray {

		resp, err := http.Get(service.Url + "actions")
		if err == nil {
			defer resp.Body.Close()
			actionsResp := ActionsResponse{}
			body, _ := io.ReadAll(resp.Body)

			if err := json.Unmarshal(body, &actionsResp); err == nil {
				actionsResponse = actionsResp.Actions
			}
		}

		resp, err = http.Get(service.Url + "reactions")
		if err == nil {
			defer resp.Body.Close()
			reactionsResp := ReactionsResponse{}
			body, _ := io.ReadAll(resp.Body)

			if err := json.Unmarshal(body, &reactionsResp); err == nil {
				reactionsResponse = reactionsResp.Reactions
			}
		}

		serviceResponse := ServiceResponse{
			Name:      service.Name,
			Actions:   actionsResponse,
			Reactions: reactionsResponse,
		}
		servicesJson = append(servicesJson, serviceResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"client": gin.H{
			"host": getIp(),
		},
		"server": gin.H{
			"current_time": time.Now().Unix(),
			"services":     servicesJson,
		},
	})
}
