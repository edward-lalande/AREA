package area_test

import (
	"bytes"
	"encoding/json"
	models "google/Models"
	"google/routes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetReactions(t *testing.T) {
	router := gin.Default()
	routes.ApplyRoutes(router)
	models.ReactionsModelPath = "../Models/Reactions.json"

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/reactions", nil)

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"name\":\"Google\",\"reactions\":[{\"arguments\":[{\"display\":\"Summary events\",\"name\":\"summary\",\"type\":\"string\"},{\"display\":\"Description events\",\"name\":\"description\",\"type\":\"string\"},{\"display\":\"Start time in format ISO-8601 date\",\"name\":\"start_time\",\"type\":\"string\"},{\"display\":\"End time in format ISO-8601 date\",\"name\":\"end_time\",\"type\":\"string\"},{\"display\":\"Attendees\",\"name\":\"attendees\",\"type\":\"string\"}],\"description\":\"Create an event in a start time end end time with attendees if you want\",\"name\":\"Create Event\",\"reaction_id\":6,\"reaction_type\":0},{\"arguments\":[{\"display\":\"Recipient\",\"name\":\"recipient\",\"type\":\"string\"},{\"display\":\"Subject\",\"name\":\"subject\",\"type\":\"string\"},{\"display\":\"Message\",\"name\":\"message\",\"type\":\"string\"}],\"description\":\"Send a mail to the recipient with the subject and the your message\",\"name\":\"Send mail\",\"reaction_id\":6,\"reaction_type\":1}]}", w.Body.String())
	models.ReactionsModelPath = "Models/Reactions.json"
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
	body := models.GoogleReaction{}
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
	assert.Equal(t, "{\"actions\":[{\"action_id\":6,\"action_type\":0,\"arguments\":[],\"description\":\"When you create an event, it's trigger\",\"name\":\"Create Event\"},{\"action_id\":6,\"action_type\":1,\"arguments\":[],\"description\":\"When you delete an event, it's trigger\",\"name\":\"Delete Event\"},{\"action_id\":6,\"action_type\":2,\"arguments\":[],\"description\":\"When you send or receive an email, it's trigger\",\"name\":\"Send/Receive new mail\"}],\"name\":\"Google\"}", w.Body.String())
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
	assert.Equal(t, "{\"error\":\"failed to connect to `user=edward database=`: /tmp/.s.PGSQL.5432 (/tmp): dial error: dial unix /tmp/.s.PGSQL.5432: connect: no such file or directory\"}{\"details\":\"invalid request\",\"error\":\"Invalid input data\"}", w.Body.String())
}

func TestStoreActionsInternalServerError(t *testing.T) {
	body := models.ReceivedActions{}
	router := gin.Default()
	routes.ApplyRoutes(router)
	jsonBody, _ := json.Marshal(body)
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/action", bytes.NewBuffer(jsonBody))

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "{\"error\":\"failed to connect to `user=edward database=`: /tmp/.s.PGSQL.5432 (/tmp): dial error: dial unix /tmp/.s.PGSQL.5432: connect: no such file or directory\"}{\"details\":\"failed to get Gmail profile: 401 Unauthorized\",\"error\":\"Invalid input data\"}", w.Body.String())
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
