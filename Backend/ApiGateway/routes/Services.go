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
	var servicesArray []serviceList
	var uppedServices []serviceList

	servicesArray = append(servicesArray, serviceList{"Date Time Services", "time", utils.GetEnvKey("TIME_API"), "#b3b3b3"})
	servicesArray = append(servicesArray, serviceList{"Discord Services", "discord", utils.GetEnvKey("DISCORD_API"), "#7289da"})

	for _, service := range servicesArray {
		_, err := http.Get(service.Url + "ping")
		if err != nil {
			continue
		}
		uppedServices = append(uppedServices, service)
	}

	c.JSON(http.StatusOK, gin.H{"services": uppedServices})
}
