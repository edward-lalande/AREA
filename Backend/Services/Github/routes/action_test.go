package routes_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	models "github/Models"
	"github/routes"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetWebhooksPush(t *testing.T) {
	router := gin.Default()
	routes.ApplyRoutes(router)

	w := httptest.NewRecorder()
	body := `{"pusher": {"name": "test_pusher"}, "commits": [{"added": ["test_value"]}]}`
	req, _ := http.NewRequest("POST", "/webhook/push", bytes.NewBufferString(body))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestGetWebhooksCommitComment(t *testing.T) {
	router := gin.Default()
	routes.ApplyRoutes(router)

	w := httptest.NewRecorder()
	body := `{"comment": {"body": "test comment"}}`
	req, _ := http.NewRequest("POST", "/webhook/commit_comment", bytes.NewBufferString(body))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetActions(t *testing.T) {
	router := gin.Default()
	routes.ApplyRoutes(router)
	models.ActionsModelsPath = "../Models/Actions.json"

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/actions", nil)

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"actions\":[{\"action_id\":4,\"action_type\":0,\"arguments\":[{\"display\":\"Pusher login\",\"name\":\"pusher\",\"type\":\"string\"}],\"description\":\"When a user push, it's trigger\",\"name\":\"On user push\"},{\"action_id\":4,\"action_type\":1,\"arguments\":[{\"display\":\"Files added count\",\"name\":\"number\",\"type\":\"number\"}],\"description\":\"When the number of file added to the commit have this count, it's trigger\",\"name\":\"On files added count\"},{\"action_id\":4,\"action_type\":2,\"arguments\":[{\"display\":\"Files modified count\",\"name\":\"number\",\"type\":\"number\"}],\"description\":\"When the number of file modified to the commit have this count, it's trigger\",\"name\":\"On files modified count\"},{\"action_id\":4,\"action_type\":3,\"arguments\":[{\"display\":\"Files deleted count\",\"name\":\"number\",\"type\":\"number\"}],\"description\":\"When the number of file deleted to the commit have this count, it's trigger\",\"name\":\"On files deleted count\"},{\"action_id\":4,\"action_type\":4,\"arguments\":[{\"display\":\"File added name\",\"name\":\"value\",\"type\":\"string\"}],\"description\":\"When the number of file added to the commit have this count, it's trigger\",\"name\":\"On files added\"},{\"action_id\":4,\"action_type\":5,\"arguments\":[{\"display\":\"File modified name\",\"name\":\"value\",\"type\":\"string\"}],\"description\":\"When the number of file modified to the commit have this count, it's trigger\",\"name\":\"On files modified\"},{\"action_id\":4,\"action_type\":6,\"arguments\":[{\"display\":\"File removed name\",\"name\":\"value\",\"type\":\"string\"}],\"description\":\"When the number of file removeds to the commit have this count, it's trigger\",\"name\":\"On files removed\"},{\"action_id\":4,\"action_type\":7,\"arguments\":[{\"display\":\"Commits count\",\"name\":\"number\",\"type\":\"number\"}],\"description\":\"When the number of the commit have this count, it's trigger\",\"name\":\"On commits count\"},{\"action_id\":4,\"action_type\":8,\"arguments\":[{\"display\":\"Commit name\",\"name\":\"value\",\"type\":\"string\"}],\"description\":\"When the commit is named to this, it's trigger\",\"name\":\"On commit named\"}],\"name\":\"Github\"}", w.Body.String())
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
}

func TestStoreActionsInternalServerError(t *testing.T) {
	body := models.GithubAction{}
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
