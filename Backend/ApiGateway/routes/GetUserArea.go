package routes

import (
	models "api-gateway/Models"
	"api-gateway/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ActionHandler func(c *gin.Context, areaId string, actionId int) map[string]interface{}

type ReactionHandler func(c *gin.Context, areaId string, reactionId int) []map[string]interface{}

var actionHandlers = map[int]map[int]ActionHandler{
	1: {0: func(c *gin.Context, areaId string, actionId int) map[string]interface{} {
		db := utils.OpenDB(c)
		if db == nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to open the database"})
			return nil
		}
		defer db.Close(c)

		var timeAction models.TimeActionDatabase
		row := db.QueryRow(c, "SELECT * FROM \"TimeAction\" WHERE area_id = $1", areaId)
		if err := row.Scan(&timeAction.Id, &timeAction.AreaId, &timeAction.Continent, &timeAction.City, &timeAction.Hour, &timeAction.Minute); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch action"})
			return nil
		}
		return map[string]interface{}{
			"action_name": "At time",
			"action_id":   actionId,
			"continent":   timeAction.Continent,
			"city":        timeAction.City,
			"hour":        timeAction.Hour,
			"minute":      timeAction.Minute,
		}
	},
	}}

var reactionHandlers = map[int]map[int]ReactionHandler{
	2: {0: func(c *gin.Context, areaId string, reactionId int) []map[string]interface{} {
		db := utils.OpenDB(c)
		if db == nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to open the database"})
			return nil
		}
		defer db.Close(c)

		var discordReaction models.DiscordReactionDatabase
		reactions := []map[string]interface{}{}

		rows, err := db.Query(c, "SELECT * FROM \"DiscordReactions\" WHERE area_id = $1", areaId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch reactions"})
			return nil
		}
		defer rows.Close()

		for rows.Next() {
			if err := rows.Scan(&discordReaction.Id, &discordReaction.AreaId, &discordReaction.ReactionType, &discordReaction.UserToken, &discordReaction.ChannelId, &discordReaction.Message); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse reaction"})
				return nil
			}
			reactions = append(reactions, map[string]interface{}{
				"reaction_name": "Send message",
				"reaction_id":   reactionId,
				"reaction_type": discordReaction.ReactionType,
				"channel_id":    discordReaction.ChannelId,
				"message":       discordReaction.Message,
			})
		}
		return reactions
	},
	}}

func getActionsTypeFromAreaId(c *gin.Context, actionId int, areaId string) int {
	if actionId == 1 {
		return 0
	}
	db := utils.OpenDB(c)
	actionType := 0
	var actions = map[int]string{
		1: "\"TimeAction\"",
		2: "\"DiscordActions\"",
	}
	if db == nil {
		return -1
	}

	defer db.Close(c)
	row := db.QueryRow(c, "SELECT action_type FROM "+actions[actionId]+" WHERE area_id = $1", areaId)
	row.Scan(&actionType)

	return actionType
}

func getReactionsTypeFromAreaId(c *gin.Context, reactionId int, areaId string) int {
	db := utils.OpenDB(c)
	reactionType := 0
	var reactions = map[int]string{
		2: "\"DiscordReactions\"",
	}
	if db == nil {
		return -1
	}
	defer db.Close(c)
	row := db.QueryRow(c, "SELECT reaction_type FROM "+reactions[reactionId]+" WHERE area_id = $1", areaId)
	row.Scan(&reactionType)

	return reactionType
}

func safeReactionHandler(c *gin.Context, reactionId int, areaId string) ([]map[string]interface{}, error) {
	reactionHandler, exists := reactionHandlers[reactionId]
	if !exists {
		return nil, fmt.Errorf("reaction handler not found for id %d", reactionId)
	}

	reactionType := getReactionsTypeFromAreaId(c, reactionId, areaId)
	if reactionType == -1 {
		return nil, fmt.Errorf("invalid reaction type for area id %s", areaId)
	}

	handler, typeExists := reactionHandler[reactionType]
	if !typeExists {
		return nil, fmt.Errorf("reaction type not found for type %d", reactionType)
	}

	return handler(c, areaId, reactionType), nil
}

func safeActionHandler(c *gin.Context, actionId int, areaId string) (map[string]interface{}, error) {
	actionHandler, exists := actionHandlers[actionId]
	if !exists {
		return nil, fmt.Errorf("action handler not found for id %d", actionId)
	}

	actionType := getActionsTypeFromAreaId(c, actionId, areaId)
	if actionType == -1 {
		return nil, fmt.Errorf("invalid action type for area id %s", areaId)
	}

	handler, typeExists := actionHandler[actionType]
	if !typeExists {
		return nil, fmt.Errorf("action type not found for type %d", actionType)
	}

	return handler(c, areaId, actionType), nil
}

// Get Area of a user
// @Summary Get area from a users in all servieces by sending an array of object of area
// @Description Get area from a users in all servieces by sending an array of object of area
// @Tags Area api-gateway
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string "Reactions name with parameters of it as object"
// @Failure 400 {object} map[string]string "Bad Requests"
// @Failure 500 {object} map[string]string "Internal error"
// @Router /areas [get]
func GetUserAreas(c *gin.Context) {
	token := c.GetHeader("token")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token is required"})
		return
	}

	db := utils.OpenDB(c)
	if db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to open the database"})
		return
	}
	defer db.Close(c)

	rows, err := db.Query(c, "SELECT * FROM \"Area\" WHERE user_token = $1", token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch areas"})
		return
	}
	defer rows.Close()

	var areas []map[string]interface{}
	for rows.Next() {
		var area models.AreaDatabase
		if err := rows.Scan(&area.Id, &area.UserToken, &area.AreaId, &area.ServiceActionId, &area.ServiceReactionId); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse area"})
			return
		}

		action, err := safeActionHandler(c, area.ServiceActionId, area.AreaId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		reactions, err := safeReactionHandler(c, area.ServiceReactionId, area.AreaId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		areas = append(areas, map[string]interface{}{
			"area_id":   area.AreaId,
			"action":    action,
			"reactions": reactions,
		})
	}

	c.JSON(http.StatusOK, gin.H{"areas": areas})
}
