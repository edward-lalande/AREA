package routes

import (
	models "api-gateway/Models"
	"api-gateway/utils"
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendTicketMasterActions(data models.TicketMasterAction, c *gin.Context) *http.Response {
	jsonBody, err := json.Marshal(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil
	}

	resp, err := http.Post(utils.GetEnvKey("TICKET_MASTER_API")+"action", "application/jsons", bytes.NewBuffer(jsonBody))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil
	}
	defer resp.Body.Close()

	return resp
}
