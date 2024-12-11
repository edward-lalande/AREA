package routes

import (
	"bytes"
	"encoding/json"
	models "message-brocker/Models"
	"message-brocker/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func findAreaInDatabase(AreaId string, c *gin.Context) int {
	db := utils.OpenDB(c)
	if db == nil {
		return -1
	}

	var serviceReactionId int
	query := `SELECT service_reaction_id FROM "Area" WHERE area_id = $1`

	err := db.QueryRow(c, query, AreaId).Scan(&serviceReactionId)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return -1
		}
		return -1
	}

	defer db.Close(c)
	return serviceReactionId
}

func Trigger(c *gin.Context) {
	var (
		receivedData models.TriggerModelGateway
		services     map[int]string = make(map[int]string)
	)
	services[0] = utils.GetEnvKey("USER_API")
	services[1] = utils.GetEnvKey("TIME_API")
	services[2] = utils.GetEnvKey("DISCORD_API")

	if err := c.ShouldBindJSON(&receivedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	serviceReactionId := findAreaInDatabase(receivedData.AreaId, c)
	if serviceReactionId == -1 || serviceReactionId > len(services) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "service Reactiond id is" + string(serviceReactionId)})
		return
	}

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(receivedData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp, err := http.Post(services[serviceReactionId]+"trigger", "application/json", &buf)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()
	c.JSON(resp.StatusCode, gin.H{"body": resp.Body})
}
