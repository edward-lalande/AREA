package area_test

import (
	models "asana/Models"
	"asana/routes"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetReactions(t *testing.T) {
	models.ReactionsPath = "../Models/Reactions.json"
	router := gin.Default()
	routes.ApplyRoutes(router)
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/reactions", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"color\":\"#ff80e1\",\"name\":\"Asana\",\"reactions\":[{\"arguments\":[{\"display\":\"Project name\",\"name\":\"project_name\",\"type\":\"string\"},{\"display\":\"Workspace id\",\"name\":\"workspace_id\",\"type\":\"string\"},{\"display\":\"Project note\",\"name\":\"note\",\"type\":\"string\"}],\"description\":\"Creation of a new project on your Asana account\",\"name\":\"Create a new Project\",\"reaction_id\":10,\"reaction_type\":0},{\"arguments\":[{\"display\":\"Task name\",\"name\":\"note\",\"type\":\"string\"},{\"display\":\"Workspace id\",\"name\":\"workspace_id\",\"type\":\"string\"},{\"display\":\"Project id\",\"name\":\"project_name\",\"type\":\"string\"}],\"description\":\"Creation of a new Task on your Asana project\",\"name\":\"Create a new Task\",\"reaction_id\":10,\"reaction_type\":1},{\"arguments\":[{\"display\":\"Project id\",\"name\":\"project_name\",\"type\":\"string\"},{\"display\":\"Mail of the member\",\"name\":\"note\",\"type\":\"string\"}],\"description\":\"Add a new member to your Asana project\",\"name\":\"Add member to project\",\"reaction_id\":10,\"reaction_type\":2}]}", w.Body.String())
	models.ReactionsPath = "Models/Reactions.json"
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
