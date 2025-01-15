package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	receivedData := models.TriggerModelGateway{}
	services := map[int]string{
		0:  utils.GetEnvKey("USER_API"),
		1:  utils.GetEnvKey("TIME_API"),
		2:  utils.GetEnvKey("DISCORD_API"),
		3:  utils.GetEnvKey("DROPBOX_API"),
		4:  utils.GetEnvKey("GITHUB_API"),
		5:  utils.GetEnvKey("GITLAB_API"),
		6:  utils.GetEnvKey("GOOGLE_API"),
		7:  utils.GetEnvKey("METEO_API"),
		9:  utils.GetEnvKey("SPOTIFY_API"),
		10: utils.GetEnvKey("ASANA_API"),
		11: utils.GetEnvKey("TICKET_MASTER_API"),
		12: utils.GetEnvKey("TWILIO_API"),
		13: utils.GetEnvKey("CRYPTOMONEY_API"),
		14: utils.GetEnvKey("MIRO_API"),
	}

	if err := c.ShouldBindJSON(&receivedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	serviceReactionId := findAreaInDatabase(receivedData.AreaId, c)
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(receivedData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	resp, err := http.Post(services[serviceReactionId]+"trigger", "application/json", &buf)

	if err != nil {
		fmt.Println("error:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	c.JSON(resp.StatusCode, gin.H{"body": resp.Body})
}
