package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DateTimeResponse struct {
	Datetime string `json:"datetime"`
}

func GetTime(c *gin.Context) {
	apiURL := "http://worldtimeapi.org/api/timezone/Etc/UTC"

	resp, err := http.Get(apiURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Impossible de contacter l'API"})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		c.JSON(resp.StatusCode, gin.H{"error": "Échec de l'appel API"})
		return
	}

	var dateTime DateTimeResponse
	if err := json.NewDecoder(resp.Body).Decode(&dateTime); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur de décodage JSON"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"date_time": dateTime.Datetime})
}
