package routes

import (
	models "api-gateway/Models"
	"api-gateway/utils"
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"log"
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

	_, err := db.Exec(context.Background(), query, userToken, areaID, serviceActionID, serviceReactionID)
	if err != nil {
		return err
	}
	defer db.Close(c)
	return nil
}

func Area(c *gin.Context) {
	var payload []models.PayloadItem
	areaID := GenerateCryptoID()

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}

	for _, item := range payload {
		var action models.BaseAction
		var reaction models.BaseReaction

		if item.Action != nil {
			if err := json.Unmarshal(*item.Action, &action); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid action data"})
				return
			}

			switch action.ActionID {
			case 1:
				var action models.TypeTimeAction
				if err := json.Unmarshal(*item.Action, &action); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Type1 action data"})
					return
				}
				resp := SendTime(areaID, action, c)
				c.JSON(http.StatusOK, gin.H{"body": resp.Body})
			}
		}

		if item.Reaction != nil {
			if err := json.Unmarshal(*item.Reaction, &reaction); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid reaction data"})
				return
			}

			switch reaction.ReactionID {
			case 2:
				var reaction models.TypeDiscordReaction
				if err := json.Unmarshal(*item.Reaction, &reaction); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Type0 reaction data"})
					return
				}
				resp := SendMessageDiscordReaction(item.UserToken, areaID, c, reaction)
				c.JSON(http.StatusOK, gin.H{"body": resp.Body})
			}
		}

		err := writeAreaInDatabase(c, areaID, item.UserToken, action.ActionID, reaction.ReactionID)
		if err != nil {
			log.Printf("Failed to write area: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to write area"})
			return
		}
	}
}
