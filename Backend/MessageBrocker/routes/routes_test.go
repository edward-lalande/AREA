package routes_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	models "message-brocker/Models"
	"message-brocker/routes"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestTrigger(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Valid request", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		payload := models.TriggerModelGateway{AreaId: "testAreaId"}
		body, _ := json.Marshal(payload)

		c.Request, _ = http.NewRequest("POST", "/trigger", bytes.NewBuffer(body))
		c.Request.Header.Set("Content-Type", "application/json")

		routes.Trigger(c)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("Invalid JSON", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		c.Request, _ = http.NewRequest("POST", "/trigger", bytes.NewBuffer([]byte("invalid json")))
		c.Request.Header.Set("Content-Type", "application/json")

		routes.Trigger(c)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}
