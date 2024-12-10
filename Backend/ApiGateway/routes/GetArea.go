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

func GetActions(c *gin.Context) {
	var servicesArray []actionsList
	var actions []map[string]interface{}

	servicesArray = append(servicesArray, actionsList{"Date Time Services", utils.GetEnvKey("TIME_API")})
	servicesArray = append(servicesArray, actionsList{"Discord Services", utils.GetEnvKey("DISCORD_API")})

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
		actions = append(actions, gin.H{service.name: jsonResponse})
	}

	c.JSON(http.StatusOK, actions)
}

func GetReactions(c *gin.Context) {
	var servicesArray []reactionsList
	var reactions []any

	servicesArray = append(servicesArray, reactionsList{"Date Time Services", utils.GetEnvKey("TIME_API")})
	servicesArray = append(servicesArray, reactionsList{"Discord Services", utils.GetEnvKey("DISCORD_API")})

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

		reactions = append(reactions, gin.H{service.name: jsonResponse})
	}

	c.JSON(http.StatusOK, reactions)
}
