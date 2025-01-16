package area_test

import (
	"bytes"
	models "cryptomoney/Models"
	"cryptomoney/routes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetActions(t *testing.T) {
	router := gin.Default()
	routes.ApplyRoutes(router)
	models.ActionPath = "../Models/Actions.json"

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/actions", nil)

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"actions\":[{\"action_id\":13,\"action_type\":0,\"arguments\":[{\"display\":\"CryptoMoney symbole\",\"name\":\"symbole\",\"type\":\"string\"},{\"display\":\"CryptoMoney Devise\",\"name\":\"devise\",\"type\":\"string\"},{\"display\":\"CryptoMoney value\",\"name\":\"value\",\"type\":\"number\"}],\"description\":\"If the CryptoMoney symbole is under the value indicate, it's trigger\",\"name\":\"CryptoMoney value is under your value\"},{\"action_id\":13,\"action_type\":1,\"arguments\":[{\"display\":\"CryptoMoney symbole\",\"name\":\"symbole\",\"type\":\"string\"},{\"display\":\"CryptoMoney Devise\",\"name\":\"devise\",\"type\":\"string\"},{\"display\":\"CryptoMoney value\",\"name\":\"value\",\"type\":\"number\"}],\"description\":\"If the CryptoMoney symbole is up to the value indicate, it's trigger\",\"name\":\"CryptoMoney value is up to your value\"},{\"action_id\":13,\"action_type\":2,\"arguments\":[{\"display\":\"CryptoMoney symbole\",\"name\":\"symbole\",\"type\":\"string\"},{\"display\":\"CryptoMoney Devise\",\"name\":\"devise\",\"type\":\"string\"},{\"display\":\"CryptoMoney value\",\"name\":\"value\",\"type\":\"number\"}],\"description\":\"If the CryptoMoney symbole is equal to the value indicate, it's trigger\",\"name\":\"CryptoMoney value is equal to your value\"}],\"name\":\"CryptoMoney\"}", w.Body.String())
	models.ActionPath = "Models/Actions.json"
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

func TestStoreActionsBadRequests(t *testing.T) {
	router := gin.Default()
	routes.ApplyRoutes(router)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/action", nil)

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "{\"details\":\"invalid request\",\"error\":\"Invalid request data\"}", w.Body.String())
}

func TestStoreActionsInternalServerError(t *testing.T) {
	body := models.Actions{
		AreaId:     "area id",
		ActionType: 0,
		Symbole:    "BTC",
		Devise:     "USD",
		Value:      12,
	}
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
