package routes

import (
	"bytes"
	models "date-time-service/Models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetActions(t *testing.T) {
	router := gin.Default()
	ApplyRoutes(router)
	models.GetActionsPath = "../Models/Actions.json"

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/actions", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"actions\":[{\"action_id\":1,\"action_type\":0,\"arguments\":[{\"display\":\"Hour\",\"name\":\"hour\",\"type\":\"number\"},{\"display\":\"Minute\",\"name\":\"minute\",\"type\":\"number\"},{\"display\":\"City\",\"name\":\"city\",\"type\":\"string\"},{\"display\":\"Continent\",\"name\":\"continent\",\"type\":\"string\"}],\"description\":\"If the it is the hour, minute in the city in the continent indicate, it's trigger\",\"name\":\"Every day at\"},{\"action_id\":1,\"action_type\":1,\"arguments\":[{\"display\":\"Minute\",\"name\":\"minute\",\"type\":\"number\"},{\"display\":\"City\",\"name\":\"city\",\"type\":\"string\"},{\"display\":\"Continent\",\"name\":\"continent\",\"type\":\"string\"}],\"description\":\"If the it is the minute in the city in the continent indicate, it's trigger\",\"name\":\"Every hour at\"},{\"action_id\":1,\"action_type\":2,\"arguments\":[{\"display\":\"Minute\",\"name\":\"minute\",\"type\":\"number\"},{\"display\":\"City\",\"name\":\"city\",\"type\":\"string\"},{\"display\":\"Continent\",\"name\":\"continent\",\"type\":\"string\"}],\"description\":\"If the it is the hour even in the minute in the city in the continent indicate, it's trigger\",\"name\":\"Hour is even\"},{\"action_id\":1,\"action_type\":3,\"arguments\":[{\"display\":\"Minute\",\"name\":\"minute\",\"type\":\"number\"},{\"display\":\"City\",\"name\":\"city\",\"type\":\"string\"},{\"display\":\"Continent\",\"name\":\"continent\",\"type\":\"string\"}],\"description\":\"If the it is the hour odd in the minute in the city in the continent indicate, it's trigger\",\"name\":\"Hour is odd\"}],\"name\":\"Date Time\"}", w.Body.String())
	models.GetActionsPath = "Models/Actions.json"
}

func TestGetActionsError(t *testing.T) {
	router := gin.Default()
	ApplyRoutes(router)
	models.GetActionsPath = "null"

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/actions", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "{\"error\":\"open null: no such file or directory\"}", w.Body.String())
	models.GetActionsPath = "Models/Actions.json"
}

func TestPostActionsBadRequest(t *testing.T) {
	router := gin.Default()
	ApplyRoutes(router)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/action", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "{\"error\":\"invalid request\"}", w.Body.String())
}

func TestPostActionsInternalError(t *testing.T) {
	router := gin.Default()
	body := models.TimeActionReceive{
		AreaId:     "id",
		ActionType: 0,
		Continent:  "",
		Hour:       1,
		Minute:     1,
		City:       "",
	}
	ApplyRoutes(router)
	jsonData, _ := json.Marshal(body)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/action", bytes.NewBuffer(jsonData))
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "{\"error\":\"Unable to open the database\"}", w.Body.String())
}
