package area_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	models "spotify/Models"
	"spotify/routes"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetReactions(t *testing.T) {
	router := gin.Default()
	routes.ApplyRoutes(router)
	models.ReactionsModelsPath = "../Models/Reactions.json"

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/reactions", nil)

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"name\":\"Spotify\",\"reactions\":[{\"arguments\":[],\"description\":\"It's pause your music on all your devices\",\"name\":\"Pause the music\",\"reaction_id\":9,\"reaction_type\":0}]}", w.Body.String())
	models.ReactionsModelsPath = "Models/Reactions.json"
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
	body := models.ReactionsReceived{}
	router := gin.Default()
	routes.ApplyRoutes(router)
	jsonBody, _ := json.Marshal(body)
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/reaction", bytes.NewBuffer(jsonBody))

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "{\"error\":\"failed to connect to `user=edward database=`: /tmp/.s.PGSQL.5432 (/tmp): dial error: dial unix /tmp/.s.PGSQL.5432: connect: no such file or directory\"}", w.Body.String())
}

func TestGetActions(t *testing.T) {
	router := gin.Default()
	routes.ApplyRoutes(router)
	models.ActionsModelsPath = "../Models/Actions.json"

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/actions", nil)

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"actions\":[{\"action_id\":9,\"action_type\":0,\"arguments\":[],\"description\":\"If you listen to a music, it's trigger\",\"name\":\"Listen to music\"},{\"action_id\":9,\"action_type\":1,\"arguments\":[],\"description\":\"If you create a new playlist, it's trigger\",\"name\":\"Create new playlist\"},{\"action_id\":9,\"action_type\":2,\"arguments\":[],\"description\":\"If you remove a playlist, it's trigger\",\"name\":\"Remove a playlist\"}],\"name\":\"Spotify\"}", w.Body.String())
	models.ActionsModelsPath = "Models/Actions.json"
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
	assert.Equal(t, "{\"error\":\"failed to connect to `user=edward database=`: /tmp/.s.PGSQL.5432 (/tmp): dial error: dial unix /tmp/.s.PGSQL.5432: connect: no such file or directory\"}{\"error\":\"invalid request\"}", w.Body.String())
}

func TestStoreActionsInternalServerError(t *testing.T) {
	body := models.ActionsData{}
	router := gin.Default()
	routes.ApplyRoutes(router)
	jsonBody, _ := json.Marshal(body)
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/action", bytes.NewBuffer(jsonBody))

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "{\"error\":\"failed to connect to `user=edward database=`: /tmp/.s.PGSQL.5432 (/tmp): dial error: dial unix /tmp/.s.PGSQL.5432: connect: no such file or directory\"}\"Unable to open the database\"", w.Body.String())
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
