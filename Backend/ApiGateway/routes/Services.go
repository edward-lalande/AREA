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
		{"Date Time Services", "time", utils.GetEnvKey("TIME_API"), "#b3b3b3"},
		{"Discord Services", "discord", utils.GetEnvKey("DISCORD_API"), "#7289da"},
		{"Dropbox", "dropbox", utils.GetEnvKey("DROPBOX_API"), "#7289da"},
		{"Github", "github", utils.GetEnvKey("GITHUB_API"), "#7289da"},
		{"Gitlab", "gitlab", utils.GetEnvKey("GITLAB_API"), "#7289da"},
		{"Google", "google", utils.GetEnvKey("GOOGLE_API"), "#7289da"},
		{"Meteo", "meteo", utils.GetEnvKey("METEO_API"), "#7289da"},
		{"Spotify", "spotify", utils.GetEnvKey("SPOTIFY_API"), "#7289da"},
		{"Asana", "asana", utils.GetEnvKey("ASANA_API"), "#7289da"},
		{"Ticket Master", "ticket-master", utils.GetEnvKey("TICKET_MASTER_API"), "#7289da"},
		{"Twilio", "twilio", utils.GetEnvKey("TWILIO_API"), "#7289da"},
	}
	var uppedServices []serviceList

	for _, service := range servicesArray {
		_, err := http.Get(service.Url + "ping")
		if err != nil {
			continue
		}
		uppedServices = append(uppedServices, service)
	}

	c.JSON(http.StatusOK, gin.H{"services": uppedServices})
}
