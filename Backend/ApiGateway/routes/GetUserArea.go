package routes

import (
	models "api-gateway/Models"
	"api-gateway/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func searchActions(c *gin.Context, actionId int, areaId string) map[string]interface{} {
	var actionsDb = map[int]string{
		1: "\"TimeAction\"",
	}
	var dbValue models.TimeActionDatabase

	db := utils.OpenDB(c)
	if db == nil {
		c.JSON(http.StatusInternalServerError, "Unable to open the database")
		return nil
	}
	defer db.Close(c)

	query, ok := actionsDb[actionId]
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid action ID"})
		return nil
	}

	row := db.QueryRow(c, "SELECT * FROM "+query+" WHERE area_id = $1", areaId)
	if err := row.Scan(&dbValue.Id, &dbValue.AreaId, &dbValue.Continent, &dbValue.City, &dbValue.Hour, &dbValue.Minute); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch action"})
		return nil
	}

	return map[string]interface{}{
		"action_id":   actionId,
		"action_type": 0,
		"continent":   dbValue.Continent,
		"city":        dbValue.City,
		"hour":        dbValue.Hour,
		"minute":      dbValue.Minute,
	}
}

func searchReactions(c *gin.Context, reactionId int, areaId string) []map[string]interface{} {
	var reactionsDb = map[int]string{
		2: "\"DiscordReactions\"",
	}
	var dbValue models.DiscordReactionDatabase
	var reactions []map[string]interface{}

	db := utils.OpenDB(c)
	if db == nil {
		c.JSON(http.StatusInternalServerError, "Unable to open the database")
		return nil
	}
	defer db.Close(c)

	query, ok := reactionsDb[reactionId]
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid reaction ID"})
		return nil
	}

	rows, err := db.Query(c, "SELECT * FROM "+query+" WHERE area_id = $1", areaId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch reactions"})
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&dbValue.Id, &dbValue.AreaId, &dbValue.ReactionType, &dbValue.UserToken, &dbValue.ChannelId, &dbValue.Message); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse reaction"})
			return nil
		}
		reactions = append(reactions, map[string]interface{}{
			"reaction_id":   reactionId,
			"reaction_type": dbValue.ReactionType,
			"channel_id":    dbValue.ChannelId,
			"message":       dbValue.Message,
		})
	}

	return reactions
}

func GetUserAreas(c *gin.Context) {
	token := c.GetHeader("token")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token is required"})
		return
	}

	db := utils.OpenDB(c)
	if db == nil {
		c.JSON(http.StatusInternalServerError, "Unable to open the database")
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

		action := searchActions(c, area.ServiceActionId, area.AreaId)

		reactions := searchReactions(c, area.ServiceReactionId, area.AreaId)

		areas = append(areas, map[string]interface{}{
			"area_id":   area.AreaId,
			"action":    action,
			"reactions": reactions,
		})
	}

	c.JSON(http.StatusOK, gin.H{"areas": areas})
}
