package routes_test

import (
	"bytes"
	models "discord-service/Models"
	"discord-service/routes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetReactions(t *testing.T) {
	router := gin.Default()
	routes.ApplyRoutes(router)
	models.ReactionModelPath = "../Models/Reactions.json"

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/reactions", nil)

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"color\":\"#7289da\",\"name\":\"Discord\",\"reactions\":[{\"arguments\":[{\"display\":\"Channel ID\",\"name\":\"channel_id\",\"type\":\"string\"},{\"display\":\"Message\",\"name\":\"message\",\"type\":\"string\"}],\"description\":\"Send a message on the channel indicated as an ID\",\"name\":\"Send message on channel\",\"reaction_id\":2,\"reaction_type\":0},{\"arguments\":[{\"display\":\"Guild ID\",\"name\":\"guild_id\",\"type\":\"string\"},{\"display\":\"Channel Name\",\"name\":\"channel_id\",\"type\":\"string\"}],\"description\":\"Create a channel on the Guild indicated as an ID\",\"name\":\"Create a text channel\",\"reaction_id\":2,\"reaction_type\":1},{\"arguments\":[{\"display\":\"Guild ID\",\"name\":\"guild_id\",\"type\":\"string\"},{\"display\":\"Channel Name\",\"name\":\"channel_id\",\"type\":\"string\"}],\"description\":\"Create a voice channel on the Guild indicated as an ID\",\"name\":\"Create a voice channel\",\"reaction_id\":2,\"reaction_type\":2},{\"arguments\":[{\"display\":\"Channel ID\",\"name\":\"channel_id\",\"type\":\"string\"}],\"description\":\"Delete a channel on the Guild indicated as an ID\",\"name\":\"Delete a channel\",\"reaction_id\":2,\"reaction_type\":3},{\"arguments\":[{\"display\":\"Channel ID\",\"name\":\"channel_id\",\"type\":\"string\"},{\"display\":\"Message ID\",\"name\":\"message\",\"type\":\"string\"}],\"description\":\"Pin a message a channel indicated as an ID\",\"name\":\"Pin a message\",\"reaction_id\":2,\"reaction_type\":4},{\"arguments\":[{\"display\":\"Channel ID\",\"name\":\"channel_id\",\"type\":\"string\"},{\"display\":\"Message ID\",\"name\":\"message\",\"type\":\"string\"}],\"description\":\"Unpin a message a channel indicated as an ID\",\"name\":\"Unpin a message\",\"reaction_id\":2,\"reaction_type\":5},{\"arguments\":[{\"display\":\"Guild ID\",\"name\":\"guild_id\",\"type\":\"string\"},{\"display\":\"Name\",\"name\":\"message\",\"type\":\"string\"}],\"description\":\"Create a role on the Guild indicated as an ID\",\"name\":\"Create a role\",\"reaction_id\":2,\"reaction_type\":6},{\"arguments\":[{\"display\":\"Guild ID\",\"name\":\"guild_id\",\"type\":\"string\"},{\"display\":\"Role ID\",\"name\":\"message\",\"type\":\"string\"}],\"description\":\"Delete a role on the Guild indicated as an ID\",\"name\":\"Delete a role\",\"reaction_id\":2,\"reaction_type\":7}]}", w.Body.String())
	models.ReactionModelPath = "Models/Reactions.json"
}

func TestGetReactionsError(t *testing.T) {
	router := gin.Default()
	routes.ApplyRoutes(router)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/reactions", nil)

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "{\"error\":\"open Models/Reactions.json: no such file or directory\"}", w.Body.String())
}

func TestStoreReactionsInternalServerError(t *testing.T) {
	body := models.ReactionReceiveData{}
	router := gin.Default()
	routes.ApplyRoutes(router)
	jsonBody, _ := json.Marshal(body)
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/reaction", bytes.NewBuffer(jsonBody))

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusInternalServerError, w.Code)

}

func TestGetActions(t *testing.T) {
	router := gin.Default()
	routes.ApplyRoutes(router)
	models.ModelPath = "../Models/Actions.json"

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/actions", nil)

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"actions\":[{\"action_id\":2,\"action_type\":0,\"arguments\":[{\"display\":\"Channel ID\",\"name\":\"channel_id\",\"type\":\"string\"},{\"display\":\"Message ID\",\"name\":\"message_id\",\"type\":\"string\"}],\"description\":\"If there is a reaction on a message in the channel, it's trigger\",\"name\":\"Reaction on message\"},{\"action_id\":2,\"action_type\":1,\"arguments\":[{\"display\":\"Channel ID\",\"name\":\"channel_id\",\"type\":\"string\"},{\"display\":\"Message ID\",\"name\":\"message_id\",\"type\":\"string\"}],\"description\":\"If there is a message pin in a the channel, it's trigger\",\"name\":\"Pin a message\"}],\"color\":\"#7289da\",\"name\":\"Discord\"}", w.Body.String())
	models.ModelPath = "Models/Actions.json"
}

func TestGetActionsError(t *testing.T) {
	router := gin.Default()
	routes.ApplyRoutes(router)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/actions", nil)

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "{\"error\":\"open Models/Actions.json: no such file or directory\"}", w.Body.String())
}

func TestStoreActionsInteralError(t *testing.T) {
	router := gin.Default()
	routes.ApplyRoutes(router)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/action", nil)

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestStoreActionsInternalServerError(t *testing.T) {
	body := models.ActiveReactionData{}
	router := gin.Default()
	routes.ApplyRoutes(router)
	jsonBody, _ := json.Marshal(body)
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/action", bytes.NewBuffer(jsonBody))

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestPingRoute(t *testing.T) {
	router := gin.Default()
	routes.ApplyRoutes(router)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"ping\":\"pong\"}", w.Body.String())
}

func TestTriggerStatusInternalServerError(t *testing.T) {
	body := models.ActiveReactionData{}

	router := gin.Default()
	routes.ApplyRoutes(router)

	b, _ := json.Marshal(body)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/trigger", bytes.NewBuffer(b))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)

}
