package routes_test

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

func TestTriggerStatusInternalServerError(t *testing.T) {
	body := models.TriggerdModels{}

	router := gin.Default()
	routes.ApplyRoutes(router)

	b, _ := json.Marshal(body)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/trigger", bytes.NewBuffer(b))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "{\"error\":\"failed to connect to `user=runner database=`: /var/run/postgresql/.s.PGSQL.5432 (/var/run/postgresql): dial error: dial unix /var/run/postgresql/.s.PGSQL.5432: connect: no such file or directory\"}{\"error\":\"Invalid request\"}", w.Body.String())
}
