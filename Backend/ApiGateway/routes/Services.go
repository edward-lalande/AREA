package routes

import (
	"api-gateway/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type serviceList struct {
	name string
	url  string
	isUp bool
}

func Services(c *gin.Context) {
	var servicesArray []serviceList
	var uppedServices []string

	servicesArray = append(servicesArray, serviceList{"User Services", utils.GetEnvKey("USER_API"), false})
	servicesArray = append(servicesArray, serviceList{"Date Time Services", utils.GetEnvKey("TIME_API"), false})

	for _, service := range servicesArray {
		_, err := http.Get(service.url + "ping")
		if err != nil {
			continue
		}
		uppedServices = append(uppedServices, service.name)
	}

	c.JSON(http.StatusOK, uppedServices)
}
