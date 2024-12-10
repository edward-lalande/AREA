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

func Services(c *gin.Context) {
	var servicesArray []serviceList
	var uppedServices []serviceList

	servicesArray = append(servicesArray, serviceList{"Date Time Services", "time", utils.GetEnvKey("TIME_API"), "#7289da"})
	servicesArray = append(servicesArray, serviceList{"Discord Services", "discord", utils.GetEnvKey("DISCORD_API"), "purple"})

	for _, service := range servicesArray {
		fmt.Println("call: ", service.Url+"ping")
		_, err := http.Get(service.Url + "ping")
		if err != nil {
			continue
		}
		fmt.Println("service append: ", service)
		uppedServices = append(uppedServices, service)
	}
	fmt.Println("upped services: ", uppedServices)

	c.JSON(http.StatusOK, gin.H{"services": uppedServices})
}
