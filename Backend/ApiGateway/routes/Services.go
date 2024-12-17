package routes

import (
	"api-gateway/utils"
	"fmt"
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
	var servicesArray []serviceList
	var uppedServices []serviceList

	servicesArray = append(servicesArray, serviceList{"Date Time Services", "time", utils.GetEnvKey("TIME_API"), "#b3b3b3"})
	servicesArray = append(servicesArray, serviceList{"Discord Services", "discord", utils.GetEnvKey("DISCORD_API"), "#7289da"})
	//servicesArray = append(servicesArray, serviceList{"Blue Sky", "blue-sky", utils.GetEnvKey("BLUESKY_API"), "#7289da"})
	//servicesArray = append(servicesArray, serviceList{"Github", "github", utils.GetEnvKey("GITHUB_API"), "#7289da"})
	//servicesArray = append(servicesArray, serviceList{"Gitlab", "gitlab", utils.GetEnvKey("GITLAB_API"), "#7289da"})
	//servicesArray = append(servicesArray, serviceList{"Google", "google", utils.GetEnvKey("GOOGLE_API"), "#7289da"})
	//servicesArray = append(servicesArray, serviceList{"Meteo", "meteo", utils.GetEnvKey("METEO_API"), "#7289da"})
	//servicesArray = append(servicesArray, serviceList{"Miro", "miro", utils.GetEnvKey("MIRO_API"), "#7289da"})
	//servicesArray = append(servicesArray, serviceList{"Spotify", "spotify", utils.GetEnvKey("SPOTIFY_API"), "#7289da"})
	//servicesArray = append(servicesArray, serviceList{"Steam", "steam", utils.GetEnvKey("STEAM_API"), "#7289da"})
	//servicesArray = append(servicesArray, serviceList{"Ticket Master", "ticket-master", utils.GetEnvKey("TICKET_MASTER_API"), "#7289da"})
	//servicesArray = append(servicesArray, serviceList{"Twilio", "twilio", utils.GetEnvKey("TWILIO_API"), "#7289da"})
	//servicesArray = append(servicesArray, serviceList{"Uber", "uber", utils.GetEnvKey("UBER_API"), "#7289da"})

	for _, service := range servicesArray {
		_, err := http.Get(service.Url + "ping")
		fmt.Println("call: ", service.Url+"ping")
		if err != nil {
			continue
		}
		uppedServices = append(uppedServices, service)
	}

	c.JSON(http.StatusOK, gin.H{"services": uppedServices})
}
