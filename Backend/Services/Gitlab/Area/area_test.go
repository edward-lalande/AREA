package area_test

import (
	"bytes"
	"encoding/json"
	models "gitlab/Models"
	"gitlab/routes"
	"net/http"
	"net/http/httptest"
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
	assert.Equal(t, "{\"name\":\"Gitlab\",\"reactions\":[{\"arguments\":[{\"display\":\"Project ID\",\"name\":\"project_id\",\"type\":\"string\"},{\"display\":\"Commentary\",\"name\":\"body\",\"type\":\"string\"}],\"name\":\"Comment last Merge Requets\",\"reaction_id\":5,\"reaction_type\":0},{\"arguments\":[{\"display\":\"Project ID\",\"name\":\"project_id\",\"type\":\"string\"},{\"display\":\"Label\",\"name\":\"body\",\"type\":\"string\"}],\"name\":\"labelise last Merge Requets\",\"reaction_id\":5,\"reaction_type\":1}]}", w.Body.String())
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
	body := models.ReceivedReactions{}
	router := gin.Default()
	routes.ApplyRoutes(router)
	jsonBody, _ := json.Marshal(body)
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/reaction", bytes.NewBuffer(jsonBody))

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "{\"error\":\"failed to connect to `user=runner database=`: /var/run/postgresql/.s.PGSQL.5432 (/var/run/postgresql): dial error: dial unix /var/run/postgresql/.s.PGSQL.5432: connect: no such file or directory\"}{\"error\":\"Invalid request\"}", w.Body.String())
}

func TestGetActions(t *testing.T) {
	router := gin.Default()
	routes.ApplyRoutes(router)
	models.ActionsModelsPath = "../Models/Actions.json"

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/actions", nil)

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"actions\":[{\"action_id\":5,\"action_type\":0,\"arguments\":[],\"description\":\"If you push, it's trigger\",\"name\":\"Push\"},{\"action_id\":5,\"action_type\":1,\"arguments\":[],\"description\":\"If you comment a Merge Requests, it's trigger\",\"name\":\"Comments\"},{\"action_id\":5,\"action_type\":2,\"arguments\":[],\"description\":\"If you create a Merge Requests, it's trigger\",\"name\":\"Merge Requests\"}],\"name\":\"Gitlab\"}", w.Body.String())
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
	assert.Equal(t, "{\"error\":\"failed to connect to `user=runner database=`: /var/run/postgresql/.s.PGSQL.5432 (/var/run/postgresql): dial error: dial unix /var/run/postgresql/.s.PGSQL.5432: connect: no such file or directory\"}{\"error\":\"Invalid request\"}", w.Body.String())
}

func TestStoreActionsInternalServerError(t *testing.T) {
	body := models.Actions{}
	router := gin.Default()
	routes.ApplyRoutes(router)
	jsonBody, _ := json.Marshal(body)
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/action", bytes.NewBuffer(jsonBody))

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "{\"error\":\"failed to connect to `user=runner database=`: /var/run/postgresql/.s.PGSQL.5432 (/var/run/postgresql): dial error: dial unix /var/run/postgresql/.s.PGSQL.5432: connect: no such file or directory\"}", w.Body.String())
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
