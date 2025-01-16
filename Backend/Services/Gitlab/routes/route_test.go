package routes_test

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

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/webhook", routes.Webhook)
	return r
}

func TestWebhookInvalidJSON(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/webhook", bytes.NewBuffer([]byte("{invalid_json")))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "\"Invalid request\"", w.Body.String())
}

func TestWebhookPushEvent(t *testing.T) {
	router := setupRouter()

	payload := map[string]interface{}{
		"object_kind": "push",
	}
	body, _ := json.Marshal(payload)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/webhook", bytes.NewBuffer(body))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestWebhookNoteEvent(t *testing.T) {
	router := setupRouter()

	payload := map[string]interface{}{
		"object_kind": "note",
	}
	body, _ := json.Marshal(payload)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/webhook", bytes.NewBuffer(body))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestWebhookMergeRequestEvent(t *testing.T) {
	router := setupRouter()

	payload := map[string]interface{}{
		"object_kind": "merge_request",
	}
	body, _ := json.Marshal(payload)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/webhook", bytes.NewBuffer(body))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestTriggerStatusInternalServerError(t *testing.T) {
	body := models.TriggerdModels{}

	router := gin.Default()
	routes.ApplyRoutes(router)

	b, _ := json.Marshal(body)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/trigger", bytes.NewBuffer(b))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)

}
