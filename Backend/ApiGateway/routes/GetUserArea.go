package routes

import (
	models "api-gateway/Models"
	"api-gateway/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetActionName(actionId int) string {
	actionsArray := []string{
		"__User",
		"TimeAction",
		"DiscordAction",
		"__Dropbox",
		"GithubActions",
		"GitlabActions",
		"GoogleActions",
		"MeteoActions",
		"__",
		"SpotifyActions",
		"__Asana",
		"TicketMasterActions",
		"__Twilio",
		"CryptoMoneyActions",
		"__Miro",
	}

	if actionsArray[actionId][0] == '_' {
		return ""
	}

	return actionsArray[actionId]
}

func GetReactionName(reactionId int) string {
	reactionsArray := []string{
		"__User",
		"__Time",
		"DiscordReactions",
		"DropboxReactions",
		"__Github",
		"GitlabReactions",
		"GoogleReactions",
		"__Meteo",
		"__",
		"SpotifyReactions",
		"AsanaReactions",
		"__TicketMaster",
		"__Twilio",
		"__CryptoMoney",
		"MiroReactions",
	}

	if reactionsArray[reactionId][0] == '_' {
		return ""
	}

	return reactionsArray[reactionId]
}

func GetAction(c *gin.Context, actionId int, areaId string) map[string]interface{} {
	db := utils.OpenDB(c)
	if db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to open the database"})
		return nil
	}
	defer db.Close(c)

	actionName := GetActionName(actionId)
	if actionName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid action name"})
		return nil
	}

	query := fmt.Sprintf("SELECT * FROM \"%s\" WHERE area_id = $1", actionName)

	rows, err := db.Query(c, query, areaId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute query"})
		return nil
	}
	defer rows.Close()

	if !rows.Next() {
		return nil
	}

	columns := rows.FieldDescriptions()
	columnNames := make([]string, len(columns))
	for i, col := range columns {
		columnNames[i] = string(col.Name)
	}

	values := make([]interface{}, len(columns))
	valuePtrs := make([]interface{}, len(columns))
	for i := range values {
		valuePtrs[i] = &values[i]
	}

	if err := rows.Scan(valuePtrs...); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan row"})
		return nil
	}

	action := make(map[string]interface{})
	for i, colName := range columnNames {
		action[colName] = values[i]
	}

	return action
}

func GetReaction(c *gin.Context, reactionId int, areaId string) map[string]interface{} {
	db := utils.OpenDB(c)
	if db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to open the database"})
		return nil
	}
	defer db.Close(c)

	reactionName := GetReactionName(reactionId)
	if reactionName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid reaction name"})
		return nil
	}

	query := fmt.Sprintf("SELECT * FROM \"%s\" WHERE area_id = $1", reactionName)

	rows, err := db.Query(c, query, areaId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute query"})
		return nil
	}
	defer rows.Close()

	if !rows.Next() {
		return nil
	}

	columns := rows.FieldDescriptions()
	columnNames := make([]string, len(columns))
	for i, col := range columns {
		columnNames[i] = string(col.Name)
	}

	values := make([]interface{}, len(columns))
	valuePtrs := make([]interface{}, len(columns))
	for i := range values {
		valuePtrs[i] = &values[i]
	}

	if err := rows.Scan(valuePtrs...); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan row"})
		return nil
	}

	reaction := make(map[string]interface{})
	for i, colName := range columnNames {
		reaction[colName] = values[i]
	}

	return reaction
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

	id := utils.ParseToken(token)
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token"})
		return
	}

	db := utils.OpenDB(c)
	if db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to open the database"})
		return
	}
	defer db.Close(c)

	rows, err := db.Query(c, "SELECT * FROM \"Area\" WHERE user_token = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch areas"})
		return
	}
	defer rows.Close()

	var areas []map[string]interface{}

	for rows.Next() {
		var area models.AreaDatabase

		if err := rows.Scan(&area.Id, &area.UserToken, &area.AreaId, &area.ServiceActionId, &area.ServiceReactionId, &area.ActionName, &area.ReactionName); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse area"})
			return
		}

		areas = append(areas, map[string]interface{}{
			"id":            area.Id,
			"area_id":       area.AreaId,
			"action_name":   area.ActionName,
			"reaction_name": area.ReactionName,
			"action":        GetAction(c, area.ServiceActionId, area.AreaId),
			"reaction":      GetReaction(c, area.ServiceReactionId, area.AreaId),
		})
	}

	c.JSON(http.StatusOK, gin.H{"areas": areas})
}

func DeleteArea(c *gin.Context) {
	var area models.AreaDatabase

	if err := c.ShouldBindJSON(&area); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}

	db := utils.OpenDB(c)
	if db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to open the database"})
		return
	}
	defer db.Close(c)

	_, err := db.Query(c, "DELETE FROM \"Area\" WHERE area_id = $1", area.AreaId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Area not found"})
		return
	}

	c.JSON(http.StatusOK, "Area deleted.")
}
