package oauth_test

import (
	"asana/routes"
	"asana/utils"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestOauthCall(t *testing.T) {
	router := gin.Default()
	routes.ApplyRoutes(router)
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/oauth", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "https://app.asana.com/-/oauth_authorize?client_id="+utils.GetEnvKey("CLIENT_ID")+
		"&redirect_uri="+utils.GetEnvKey("REDIRECT_URI")+
		"&response_type=code"+
		"&state=default"+
		"&scope=default"+
		"&grant_type=authorization_code", w.Body.String())
}

func TestAddOauthCall(t *testing.T) {
	router := gin.Default()
	routes.ApplyRoutes(router)
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/oauth", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "https://app.asana.com/-/oauth_authorize?client_id="+utils.GetEnvKey("CLIENT_ID")+
		"&redirect_uri="+utils.GetEnvKey("REDIRECT_URI_ADD")+
		"&response_type=code"+
		"&state=default"+
		"&scope=default"+
		"&grant_type=authorization_code", w.Body.String())
}

func TestCallBackStatusBadRequest(t *testing.T) {
	router := gin.Default()
	routes.ApplyRoutes(router)
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/callback", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "{\"error\":\"Missing code\"}", w.Body.String())
}

func TestAddCallBack(t *testing.T) {
	router := gin.Default()
	routes.ApplyRoutes(router)
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/add-callback?code=123", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusFound, w.Code)
	assert.Equal(t, "<a href=\"http://localhost:8081/account?asana_code=123\">Found</a>.\n\n", w.Body.String())
}

func TestAddCallBackStatusBadRequest(t *testing.T) {
	router := gin.Default()
	routes.ApplyRoutes(router)
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/add-callback", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "{\"error\":\"Missing code\"}", w.Body.String())
}

func TestGetAccessTokenBadRequest(t *testing.T) {
	body := struct {
		Test string `json:"test"`
	}{
		Test: "okay okay",
	}
	router := gin.Default()
	routes.ApplyRoutes(router)
	w := httptest.NewRecorder()

	b, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "/access-token", bytes.NewBuffer(b))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestAddAccessTokenBadRequest(t *testing.T) {
	body := struct {
		Test string `json:"test"`
	}{
		Test: "okay okay",
	}
	router := gin.Default()
	routes.ApplyRoutes(router)
	w := httptest.NewRecorder()

	b, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "/add-access-token", bytes.NewBuffer(b))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
