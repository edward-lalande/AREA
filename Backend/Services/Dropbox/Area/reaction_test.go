package area_test

import (
	"bytes"
	area "dropbox/Area"
	models "dropbox/Models"
	"dropbox/routes"
	"encoding/json"
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
	assert.Equal(t, "{\"color\":\"#0061FE\",\"name\":\"Dropbox\",\"reactions\":[{\"arguments\":[{\"display\":\"From file\",\"name\":\"from_path\",\"type\":\"string\"},{\"display\":\"To file\",\"name\":\"to_path\",\"type\":\"string\"}],\"description\":\"Rename File as the file name indicated that must begin from /\",\"name\":\"Rename File\",\"reaction_id\":3,\"reaction_type\":0},{\"arguments\":[{\"display\":\"File to share\",\"name\":\"filepath_share\",\"type\":\"string\"}],\"description\":\"Share a file to your group that must begin from /\",\"name\":\"Share File\",\"reaction_id\":3,\"reaction_type\":1}]}", w.Body.String())
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

func TestIsSlash(t *testing.T) {
	assert.Equal(t, "/hello", area.IsSlash("hello"))
	assert.Equal(t, "/hello", area.IsSlash("/hello"))
}

func TestStoreReactionsInternalServerError(t *testing.T) {
	body := models.DropBoxReactions{}
	router := gin.Default()
	routes.ApplyRoutes(router)
	jsonBody, _ := json.Marshal(body)
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/reaction", bytes.NewBuffer(jsonBody))

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
