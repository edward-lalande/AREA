package routes

import (
	models "api-gateway/Models"
	"api-gateway/utils"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GenerateCryptoID() string {
	bytes := make([]byte, 128)
	if _, err := rand.Read(bytes); err != nil {
		panic(err)
	}
	return hex.EncodeToString(bytes)
}

func writeAreaInDatabase(c *gin.Context, areaID, userToken string, serviceActionID int, serviceReactionID int, actionName string, reactionName string) error {
	id := utils.ParseToken(userToken)

	if id == "" {
		fmt.Println("no id")
		return nil
	}

	fmt.Println(actionName)
	fmt.Println(reactionName)

	query := `
		INSERT INTO "Area" (user_token, area_id, service_action_id, service_reaction_id, action_name, reaction_name)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id;
	`
	db := utils.OpenDB(c)
	if db == nil {
		return nil
	}
	_, err := db.Exec(c, query, id, areaID, serviceActionID, serviceReactionID, actionName, reactionName)
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
			case 2:
				var actionData models.TypeDiscordAction
				if err := json.Unmarshal(*item.Action, &actionData); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Type action data"})
					return
				}
				resp := sendDiscordAction(action.UserToken, areaID, c, actionData)
				c.JSON(http.StatusOK, gin.H{"body": resp.Body})
			case 4:
				var actionData models.TypeGithubAction
				if err := json.Unmarshal(*item.Action, &actionData); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Type1 action data"})
					return
				}
				resp := sendGithub(areaID, item.UserToken, c, actionData)
				c.JSON(http.StatusOK, gin.H{"body": resp.Body})
			case 5:
				var actionData models.GitlabAction
				if err := json.Unmarshal(*item.Action, &actionData); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Type5 action data"})
					return
				}
				resp := SendGitlab(areaID, actionData, c)
				c.JSON(http.StatusOK, gin.H{"body": resp.Body})
			case 6:
				var actionData models.GoogleAction
				if err := json.Unmarshal(*item.Action, &actionData); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Type6 action data"})
					return
				}
				actionData.UserToken = item.UserToken
				resp := SendGoogle(areaID, actionData, c)
				c.JSON(http.StatusOK, gin.H{"body": resp.Body})
			case 7:
				var actionData models.MeteoActions
				if err := json.Unmarshal(*item.Action, &actionData); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Type7 action data"})
					return
				}
				resp := SendMeteo(areaID, actionData, c)
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
			case 11:
				var actionData models.TicketMasterAction
				if err := json.Unmarshal(*item.Action, &actionData); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Type11 action data"})
					return
				}
				actionData.AreaID = areaID
				resp := SendTicketMasterActions(actionData, c)
				c.JSON(http.StatusOK, gin.H{"body": resp.Body})
			case 13:
				var actionData models.CryptoMoneyActions
				if err := json.Unmarshal(*item.Action, &actionData); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Type11 action data"})
					return
				}
				actionData.AreaId = areaID
				resp := SendCryptoMoneyActions(actionData, c)
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
			case 3:
				var reactionDetail models.DropBoxReactions
				if err := json.Unmarshal(*reactionData, &reactionDetail); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return
				}
				reactionDetail.AreaId = areaID
				reactionDetail.UserToken = item.UserToken
				resp := SendMessageDropbox(c, reactionDetail)
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
			case 6:
				var actionData models.GoogleReaction
				if err := json.Unmarshal(*reactionData, &actionData); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Type6 action data"})
					return
				}
				actionData.UserToken = item.UserToken
				resp := SendGoogleReactions(areaID, actionData, c)
				c.JSON(http.StatusOK, gin.H{"body": resp.Body})
			case 5:
				var reactionDetail models.GitlabReactions
				if err := json.Unmarshal(*reactionData, &reactionDetail); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return
				}
				reactionDetail.AreaId = areaID
				reactionDetail.UserToken = item.UserToken
				resp := SendGitlabReaction(reactionDetail, c)
				c.JSON(http.StatusOK, gin.H{"body": resp.Body})
			case 10:
				var reactionDetail models.AsanaReactions
				if err := json.Unmarshal(*reactionData, &reactionDetail); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return
				}
				reactionDetail.AreaId = areaID
				reactionDetail.UserToken = item.UserToken
				resp := SendAsanaReaction(reactionDetail, c)
				c.JSON(http.StatusOK, gin.H{"body": resp.Body})
			case 14:
				var reactionDetail models.MiroReactions
				if err := json.Unmarshal(*reactionData, &reactionDetail); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return
				}
				reactionDetail.AreaId = areaID
				reactionDetail.UserToken = item.UserToken
				resp := SendMiroReaction(reactionDetail, c)
				c.JSON(http.StatusOK, gin.H{"body": resp.Body})
			}
		}

		fmt.Println("step forward")

		for _, reactionData := range item.Reactions {
			var reaction models.BaseReaction
			_ = json.Unmarshal(*reactionData, &reaction)

			fmt.Println("go write area in db")

			err := writeAreaInDatabase(c, areaID, item.UserToken, action.ActionID, reaction.ReactionID, action.Name, reaction.Name)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to write area"})
				return
			}
		}
	}
}
