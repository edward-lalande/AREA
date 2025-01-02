package routes

import (
	"bytes"
	"encoding/json"
	models "gitlab/Models"
	"gitlab/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func areaIdFromActionType(c *gin.Context, actionType int) string {
	areaId := ""
	db := utils.OpenDB(c)

	query := `SELECT area_id FROM "GitlabActions" WHERE action_type = $1`

	err := db.QueryRow(c, query, actionType).Scan(&areaId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve area_id"})
		return ""
	}

	defer db.Close(c)
	return areaId
}

func sendTrigger(areaId string) {
	send := models.TriggerdModels{areaId}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(send); err != nil {
		return
	}
	http.Post(utils.GetEnvKey("MESSAGE_BROCKER")+"trigger", "application/json", &buf)
}

// GetWebhooks processes Gitlab events
// @Tags Gitlab Webhook
// @Summary Processes Gitlab events
// @Description Handles incoming webhook events and triggers actions
// @Param data body models.Webhoo true "Webhook data"
// @Success 200 {object} string "Webhook processed successfully"
// @Failure 400 {object} string "Invalid JSON payload"
// @Router /webhook [post]
func Webhook(c *gin.Context) {
	var receivedData map[string]interface{} = make(map[string]interface{})

	if err := c.ShouldBindJSON(&receivedData); err != nil {
		c.JSON(http.StatusBadRequest, "Invalid request")
		return
	}
	switch receivedData["object_kind"] {
	case "push":
		sendTrigger(areaIdFromActionType(c, 0))
	case "note":
		sendTrigger(areaIdFromActionType(c, 1))
	case "merge_request":
		sendTrigger(areaIdFromActionType(c, 2))
	}
}
