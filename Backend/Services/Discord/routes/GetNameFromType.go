package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetActionsName(c *gin.Context) {
	var actions map[int]string = make(map[int]string)
	actions[0] = "Reaction on message"
	actionType := c.Query("type")

	value, _ := strconv.Atoi(actionType)

	if value > len(actions) {
		c.String(http.StatusInternalServerError, "Actions Type didn't exists")
		return
	}

	c.String(http.StatusOK, actions[value])
}

func GetReactionsName(c *gin.Context) {
	var reactions map[int]string = make(map[int]string)
	reactions[0] = "Send Message"
	reactionType := c.Query("type")

	value, _ := strconv.Atoi(reactionType)

	if value > len(reactions) {
		c.String(http.StatusInternalServerError, "Reactions Type didn't exists")
		return
	}

	c.String(http.StatusOK, reactions[value])
}
