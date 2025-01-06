package routes

import (
	"api-gateway/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type serviceList struct {
	Name             string `json:"name"`
	CallToApiGateway string `json:"call_to_api_gateway"`
	Url              string `json:"url"`
	Color            string `json:"color"`
}

type serviceSendList struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}

// Get Up services
// @Summary Get all services up
// @Description Get services up with name to display, routes to call it to the api-gateway, url and color to display
// @Tags Area api-gateway
// @Accept json
// @Produce json
// @Success 200 {object} serviceList "services up with name to display, routes to call it to the api-gateway, url and color to display"
// @Failure 500 {object} map[string]string "Internal error"
// @Router /services [get]
func Services(c *gin.Context) {
	servicesArray := []serviceList{
		{"Date Time Service", "time", utils.GetEnvKey("TIME_API"), "#b3b3b3"},
		{"Discord Service", "discord", utils.GetEnvKey("DISCORD_API"), "#7289da"},
		{"Dropbox Service", "dropbox", utils.GetEnvKey("DROPBOX_API"), "#0061FE"},
		{"Github Service", "github", utils.GetEnvKey("GITHUB_API"), "black"},
		{"Gitlab Service", "gitlab", utils.GetEnvKey("GITLAB_API"), "#fc6d26"},
		{"Google Service", "google", utils.GetEnvKey("GOOGLE_API"), "#0F9D58"},
		{"Meteo Service", "meteo", utils.GetEnvKey("METEO_API"), "#4285F4"},
		{"Spotify Service", "spotify", utils.GetEnvKey("SPOTIFY_API"), "#1db954"},
		{"Asana Service", "asana", utils.GetEnvKey("ASANA_API"), "#ff80e1"},
		{"Ticket Master Service", "ticket-master", utils.GetEnvKey("TICKET_MASTER_API"), "#1c24ff"},
		{"Twilio Service", "twilio", utils.GetEnvKey("TWILIO_API"), "#c72e61"},
		{"Cryptomoney Service", "Cryptomoney", utils.GetEnvKey("CRYPTOMONEY_API"), "#f7931a"},
		{"Miro Service", "Miro", utils.GetEnvKey("MIRO_API"), "#faca00"},
	}
	var uppedServices []serviceSendList

	for _, service := range servicesArray {
		_, err := http.Get(service.Url + "ping")
		if err != nil {
			continue
		}
		uppedServices = append(uppedServices, serviceSendList{service.Name, service.Color})
	}

	c.JSON(http.StatusOK, gin.H{"services": uppedServices})
}
