package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"slices"

	models "github/Models"
	"github/utils"

	"github.com/gin-gonic/gin"
)

// SendTrigger sends a trigger to the message broker
// @Summary Sends a trigger to the message broker
// @Description Sends a trigger based on the provided Area ID
// @Param areaId path string true "Area ID"
// @Success 200 {object} string "Trigger sent successfully"
// @Failure 400 {object} string "Failed to encode the payload"
// @Router /trigger [post]
func SendTrigger(areaId string) {
	var buf bytes.Buffer

	send := models.GithubSendReaction{}
	send.ReactionId = areaId

	if err := json.NewEncoder(&buf).Encode(send); err != nil {
		return
	}
	http.Post(utils.GetEnvKey("MESSAGE_BROCKER")+"trigger", "application/json", &buf)
}

// CheckWebhooksPush checks webhook push data against user-defined actions
// @Summary Checks webhook push data
// @Description Matches the incoming webhook push data with user-defined GitHub actions
// @Param user body models.GithubAction true "User GitHub Action"
// @Param data body models.WebhookPush true "Webhook push data"
// @Success 200 {object} string "Trigger processed successfully"
// @Failure 400 {object} string "Failed to process the trigger"
// @Router /webhook/push/check [post]
func CheckWebhooksPush(user models.GithubAction, data models.WebhookPush) {

	if user.ActionType == 0 && user.Pusher == data.Pusher.Name {
		SendTrigger(user.AreaId)
	}

	filesAdded := 0
	filedModified := 0
	filesDeleted := 0

	for _, commit := range data.Commits {
		filesAdded += len(commit.Added)
		filedModified += len(commit.Modified)
		filesDeleted += len(commit.Removed)

		if user.ActionType == 4 && slices.Contains(commit.Added, user.Value) {
			SendTrigger(user.AreaId)
		}

		if user.ActionType == 5 && slices.Contains(commit.Modified, user.Value) {
			SendTrigger(user.AreaId)
		}

		if user.ActionType == 6 && slices.Contains(commit.Removed, user.Value) {
			SendTrigger(user.AreaId)
		}

		if user.ActionType == 8 && commit.Message == user.Value {
			SendTrigger(user.AreaId)
		}
	}

	if user.ActionType == 1 && user.Number == filesAdded {
		SendTrigger(user.AreaId)
	}

	if user.ActionType == 2 && user.Number == filedModified {
		SendTrigger(user.AreaId)
	}

	if user.ActionType == 3 && user.Number == filesDeleted {
		SendTrigger(user.AreaId)
	}

	if user.ActionType == 7 && user.Number == len(data.Commits) {
		SendTrigger(user.AreaId)
	}

}

// GetWebhooksPush processes GitHub push events
// @Summary Processes GitHub push events
// @Description Handles incoming webhook push events and triggers actions
// @Param data body models.WebhookPush true "Webhook push data"
// @Success 200 {object} string "Webhook processed successfully"
// @Failure 400 {object} string "Invalid JSON payload"
// @Router /webhook/push [post]
func GetWebhooksPush(c *gin.Context) {

	var (
		user models.GithubAction
		data models.WebhookPush
	)

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(data)

	db := utils.OpenDB(c)
	rows, err := db.Query(c, "SELECT area_id, action_type, user_token, pusher, value, number FROM \"GithubActions\" WHERE action_type < '9'")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch actions"})
		return
	}
	defer rows.Close()

	for rows.Next() {

		if err := rows.Scan(&user.AreaId, &user.ActionType, &user.UserToken, &user.Pusher, &user.Value, &user.Number); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse action"})
		}

		CheckWebhooksPush(user, data)

	}

	defer db.Close(c)
}

// GetWebhooksCommitComment processes commit comment events
// @Summary Processes GitHub commit comment events
// @Description Handles incoming webhook commit comment events
// @Param data body models.WebhooksCommitComment true "Webhook commit comment data"
// @Success 200 {object} string "Commit comment processed successfully"
// @Failure 400 {object} string "Invalid JSON payload"
// @Router /webhook/commit_comment [post]
func GetWebhooksCommitComment(c *gin.Context) {

	var (
		data models.WebhooksCommitComment
	)

	fmt.Println("Commit comment received!")

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(data)
}
