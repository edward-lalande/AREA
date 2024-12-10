package routes

import (
	"api-gateway/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type serviceList struct {
	name  string
	url   string
	color string
}

func Services(c *gin.Context) {
	var servicesArray []serviceList
	var uppedServices []serviceList

	servicesArray = append(servicesArray, serviceList{"Date Time Services", utils.GetEnvKey("TIME_API"), "white"})
	servicesArray = append(servicesArray, serviceList{"Discord Services", utils.GetEnvKey("DISCORD_API"), "purple"})

	for _, service := range servicesArray {
		_, err := http.Get(service.url + "ping")
		if err != nil {
			continue
		}
		uppedServices = append(uppedServices, service)
	}

	c.JSON(http.StatusOK, uppedServices)
}
