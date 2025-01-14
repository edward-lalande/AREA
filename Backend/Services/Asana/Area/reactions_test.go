package area_test

import (
	"asana/routes"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestTriggerStatusBadRequest(t *testing.T) {
	body := struct {
		Test int `json:"test"`
	}{
		Test: 1,
	}
	router := gin.Default()
	routes.ApplyRoutes(router)
	w := httptest.NewRecorder()

	b, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "/trigger", bytes.NewBuffer(b))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestStoreReactionsStatusBadRequest(t *testing.T) {
	body := struct {
		Test int `json:"test"`
	}{
		Test: 1,
	}
	router := gin.Default()
	routes.ApplyRoutes(router)
	w := httptest.NewRecorder()

	b, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "/reaction", bytes.NewBuffer(b))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}
