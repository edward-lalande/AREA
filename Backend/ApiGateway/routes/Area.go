package routes

import (
	models "api-gateway/Models"
	"api-gateway/utils"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GenerateCryptoID() string {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		panic(err)
	}
	return hex.EncodeToString(bytes)
}

func writeAreaInDatabase(c *gin.Context, areaID, userToken string, serviceActionID int, serviceReactionID int) error {
	query := `
		INSERT INTO "Area" (user_token, area_id, service_action_id, service_reaction_id)
		VALUES ($1, $2, $3, $4)
		RETURNING id;
	`
	db := utils.OpenDB(c)
	if db == nil {
		return nil
	}

	_, err := db.Exec(c, query, userToken, areaID, serviceActionID, serviceReactionID)
	if err != nil {
		return err
	}
	defer db.Close(c)
	return nil
}

// Create Area doc
// @Summary Create a new actions-reactions or actions-multiple reactions
// @Description Create a new combination of action and reaction (Area) for a users
// @Tags Area api-gateway
// @Accept json
// @Produce json
// @Param payload body []models.PayloadItem true "Data for all actions-reactions"
// @Success 200 {object} map[string]string "Response of all Services with the details of the executions"
// @Failure 400 {object} map[string]string "Invalid request"
// @Failure 500 {object} map[string]string "Internal error"
// @Router /areas [post]
func Area(c *gin.Context) {
	var payload []models.PayloadItem
	areaID := GenerateCryptoID()

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}

	for _, item := range payload {
		var action models.BaseAction

		if item.Action != nil {
			if err := json.Unmarshal(*item.Action, &action); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid action data"})
				return
			}

			switch action.ActionID {
			case 1:
				var actionData models.TypeTimeAction
				if err := json.Unmarshal(*item.Action, &actionData); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Type1 action data"})
					return
				}
				resp := SendTime(areaID, actionData, c)
				c.JSON(http.StatusOK, gin.H{"body": resp.Body})
			case 9:
				var actionData models.SpotifyActions
				if err := json.Unmarshal(*item.Action, &actionData); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Type9 action data"})
					return
				}
				actionData.AreaId = areaID
				resp := SendSpotifyActions(actionData, c)
				c.JSON(http.StatusOK, gin.H{"body": resp.Body})
			}
		}

		for _, reactionData := range item.Reactions {
			var reaction models.BaseReaction
			if err := json.Unmarshal(*reactionData, &reaction); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			switch reaction.ReactionID {
			case 2:
				var reactionDetail models.TypeDiscordReaction
				if err := json.Unmarshal(*reactionData, &reactionDetail); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return
				}
				resp := SendMessageDiscordReaction(item.UserToken, areaID, c, reactionDetail)
				c.JSON(http.StatusOK, gin.H{"body": resp.Body})
			case 9:
				var reactionDetail models.SpotifyReactions
				if err := json.Unmarshal(*reactionData, &reactionDetail); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return
				}
				reactionDetail.AreaId = areaID
				resp := SendSpotifyReactions(reactionDetail, c)
				c.JSON(http.StatusOK, gin.H{"body": resp.Body})
			}
		}

		for _, reactionData := range item.Reactions {
			var reaction models.BaseReaction
			_ = json.Unmarshal(*reactionData, &reaction)

			err := writeAreaInDatabase(c, areaID, item.UserToken, action.ActionID, reaction.ReactionID)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to write area"})
				return
			}
		}
	}
}
