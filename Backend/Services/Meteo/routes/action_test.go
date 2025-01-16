package routes

import (
	"bytes"
	"encoding/json"
	models "meteo/Models"
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
	assert.Equal(t, "{\"actions\":[{\"action_id\":7,\"action_type\":0,\"arguments\":[{\"display\":\"Temperature in degrees celsius\",\"name\":\"value\",\"type\":\"number\"},{\"display\":\"Latitude localisation\",\"name\":\"latitude\",\"type\":\"string\"},{\"display\":\"Longitude localisation\",\"name\":\"longitude\",\"type\":\"string\"}],\"description\":\"If in the day there is this temperature degrees celsius at the location in latitude and longitude\",\"name\":\"At temperature in the day\"},{\"action_id\":7,\"action_type\":1,\"arguments\":[{\"display\":\"Wind speed in km/h\",\"name\":\"value\",\"type\":\"number\"},{\"display\":\"Latitude localisation\",\"name\":\"latitude\",\"type\":\"string\"},{\"display\":\"Longitude localisation\",\"name\":\"longitude\",\"type\":\"string\"}],\"description\":\"If in the day the wind speed is at this in km/h at the location in latitude and longitude\",\"name\":\"At wind speed in the day\"},{\"action_id\":7,\"action_type\":2,\"arguments\":[{\"display\":\"Humidity in %\",\"name\":\"value\",\"type\":\"number\"},{\"display\":\"Latitude localisation\",\"name\":\"latitude\",\"type\":\"string\"},{\"display\":\"Longitude localisation\",\"name\":\"longitude\",\"type\":\"string\"}],\"description\":\"If in the day the humidity in the day is at this in % at the location in latitude and longitude\",\"name\":\"At humidity in the day\"}],\"name\":\"Météo\"}", w.Body.String())
	models.GetActionsPath = "Models/Actions.json"
}

func TestPostActions(t *testing.T) {
	body := models.MeteoActions{
		ActionType: 0,
		AreaId:     "id",
		Latitude:   "2",
		Longitude:  "48",
		Value:      10,
	}
	router := gin.Default()
	ApplyRoutes(router)
	w := httptest.NewRecorder()
	b, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "/action", bytes.NewBuffer(b))

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusInternalServerError, w.Code)

}

func TestPostActionsOne(t *testing.T) {
	body := models.MeteoActions{
		ActionType: 1,
		AreaId:     "id",
		Latitude:   "2",
		Longitude:  "48",
		Value:      10,
	}
	router := gin.Default()
	ApplyRoutes(router)
	w := httptest.NewRecorder()
	b, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "/action", bytes.NewBuffer(b))

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusInternalServerError, w.Code)

}

func TestPostActionsTwo(t *testing.T) {
	body := models.MeteoActions{
		ActionType: 2,
		AreaId:     "id",
		Latitude:   "2",
		Longitude:  "48",
		Value:      10,
	}
	router := gin.Default()
	ApplyRoutes(router)
	w := httptest.NewRecorder()
	b, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "/action", bytes.NewBuffer(b))

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusInternalServerError, w.Code)

}

func TestPingRoute(t *testing.T) {
	router := gin.Default()
	ApplyRoutes(router)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"ping\":\"pong\"}", w.Body.String())
}
