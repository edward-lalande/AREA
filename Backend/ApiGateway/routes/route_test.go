package routes_test

import (
	models "api-gateway/Models"
	"api-gateway/routes"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAboutOutput(t *testing.T) {
	router := gin.Default()
	routes.ApplyRoutes(router)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/about.json", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "client")
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

func TestOauth(t *testing.T) {
	router := gin.Default()
	routes.ApplyRoutes(router)
	tests := []struct {
		Link              string
		CodeExcepected    int
		ResponseExcpected string
	}{
		{
			Link:              "/discord/oauth",
			CodeExcepected:    400,
			ResponseExcpected: "{\"error\":\"Get \\\"oauth\\\": unsupported protocol scheme \\\"\\\"\"}",
		},
		{
			Link:              "/spotify/oauth",
			CodeExcepected:    400,
			ResponseExcpected: "{\"error\":\"Get \\\"oauth\\\": unsupported protocol scheme \\\"\\\"\"}",
		},
		{
			Link:              "/github/oauth",
			CodeExcepected:    400,
			ResponseExcpected: "{\"error\":\"Get \\\"oauth\\\": unsupported protocol scheme \\\"\\\"\"}",
		},
		{
			Link:              "/gitlab/oauth",
			CodeExcepected:    400,
			ResponseExcpected: "{\"error\":\"Get \\\"oauth\\\": unsupported protocol scheme \\\"\\\"\"}",
		},
		{
			Link:              "/google/oauth",
			CodeExcepected:    400,
			ResponseExcpected: "{\"error\":\"Get \\\"oauth\\\": unsupported protocol scheme \\\"\\\"\"}",
		},
		{
			Link:              "/dropbox/oauth",
			CodeExcepected:    400,
			ResponseExcpected: "{\"error\":\"Get \\\"oauth\\\": unsupported protocol scheme \\\"\\\"\"}",
		},
		{
			Link:              "/asana/oauth",
			CodeExcepected:    400,
			ResponseExcpected: "{\"error\":\"Get \\\"oauth\\\": unsupported protocol scheme \\\"\\\"\"}",
		},
		{
			Link:              "/miro/oauth",
			CodeExcepected:    400,
			ResponseExcpected: "{\"error\":\"Get \\\"oauth\\\": unsupported protocol scheme \\\"\\\"\"}",
		},
	}

	w := httptest.NewRecorder()

	for _, test := range tests {
		req, _ := http.NewRequest("GET", test.Link, nil)
		router.ServeHTTP(w, req)
		assert.Equal(t, w.Code, test.CodeExcepected)
		assert.Contains(t, w.Body.String(), test.ResponseExcpected)
	}
}

func TestAddOauth(t *testing.T) {
	router := gin.Default()
	routes.ApplyRoutes(router)
	tests := []struct {
		Link              string
		CodeExcepected    int
		ResponseExcpected string
	}{
		{
			Link:              "/discord/add-oauth",
			CodeExcepected:    400,
			ResponseExcpected: "{\"error\":\"Get \\\"add-oauth\\\": unsupported protocol scheme \\\"\\\"\"}",
		},
		{
			Link:              "/spotify/add-oauth",
			CodeExcepected:    400,
			ResponseExcpected: "{\"error\":\"Get \\\"add-oauth\\\": unsupported protocol scheme \\\"\\\"\"}",
		},
		{
			Link:              "/github/add-oauth",
			CodeExcepected:    400,
			ResponseExcpected: "{\"error\":\"Get \\\"add-oauth\\\": unsupported protocol scheme \\\"\\\"\"}",
		},
		{
			Link:              "/gitlab/add-oauth",
			CodeExcepected:    400,
			ResponseExcpected: "{\"error\":\"Get \\\"add-oauth\\\": unsupported protocol scheme \\\"\\\"\"}",
		},
		{
			Link:              "/google/add-oauth",
			CodeExcepected:    400,
			ResponseExcpected: "{\"error\":\"Get \\\"add-oauth\\\": unsupported protocol scheme \\\"\\\"\"}",
		},
		{
			Link:              "/dropbox/add-oauth",
			CodeExcepected:    400,
			ResponseExcpected: "{\"error\":\"Get \\\"add-oauth\\\": unsupported protocol scheme \\\"\\\"\"}",
		},
		{
			Link:              "/asana/add-oauth",
			CodeExcepected:    400,
			ResponseExcpected: "{\"error\":\"Get \\\"add-oauth\\\": unsupported protocol scheme \\\"\\\"\"}",
		},
		{
			Link:              "/miro/add-oauth",
			CodeExcepected:    400,
			ResponseExcpected: "{\"error\":\"Get \\\"add-oauth\\\": unsupported protocol scheme \\\"\\\"\"}",
		},
	}

	w := httptest.NewRecorder()

	for _, test := range tests {
		req, _ := http.NewRequest("GET", test.Link, nil)
		router.ServeHTTP(w, req)
		assert.Equal(t, w.Code, test.CodeExcepected)
		assert.Contains(t, w.Body.String(), test.ResponseExcpected)
	}
}

func TestGetAccessToken(t *testing.T) {
	tests := []struct {
		Oauth        models.OauthCode
		Link         string
		CodeExpected int
	}{
		{
			Oauth:        models.OauthCode{},
			Link:         "/discord/access-token",
			CodeExpected: 400,
		},
		{
			Oauth:        models.OauthCode{},
			Link:         "/spotify/access-token",
			CodeExpected: 400,
		},
		{
			Oauth:        models.OauthCode{},
			Link:         "/github/access-token",
			CodeExpected: 400,
		},
		{
			Oauth:        models.OauthCode{},
			Link:         "/gitlab/access-token",
			CodeExpected: 400,
		},
		{
			Oauth:        models.OauthCode{},
			Link:         "/google/access-token",
			CodeExpected: 400,
		},
		{
			Oauth:        models.OauthCode{},
			Link:         "/dropbox/access-token",
			CodeExpected: 400,
		},
		{
			Oauth:        models.OauthCode{},
			Link:         "/asana/access-token",
			CodeExpected: 400,
		},
		{
			Oauth:        models.OauthCode{},
			Link:         "/miro/access-token",
			CodeExpected: 400,
		},
	}
	router := gin.Default()
	routes.ApplyRoutes(router)
	w := httptest.NewRecorder()

	for _, test := range tests {

		b, _ := json.Marshal(test.Oauth)
		req, _ := http.NewRequest("POST", test.Link, bytes.NewBuffer(b))
		router.ServeHTTP(w, req)
		assert.Equal(t, test.CodeExpected, w.Code)
	}
}

func TestAddAccessToken(t *testing.T) {
	tests := []struct {
		Oauth        models.OauthCode
		Link         string
		CodeExpected int
	}{
		{
			Oauth:        models.OauthCode{},
			Link:         "/discord/access-token",
			CodeExpected: 400,
		},
		{
			Oauth:        models.OauthCode{},
			Link:         "/spotify/add-access-token",
			CodeExpected: 400,
		},
		{
			Oauth:        models.OauthCode{},
			Link:         "/github/add-access-token",
			CodeExpected: 400,
		},
		{
			Oauth:        models.OauthCode{},
			Link:         "/gitlab/add-access-token",
			CodeExpected: 400,
		},
		{
			Oauth:        models.OauthCode{},
			Link:         "/google/add-access-token",
			CodeExpected: 400,
		},
		{
			Oauth:        models.OauthCode{},
			Link:         "/dropbox/add-access-token",
			CodeExpected: 400,
		},
		{
			Oauth:        models.OauthCode{},
			Link:         "/asana/add-access-token",
			CodeExpected: 400,
		},
		{
			Oauth:        models.OauthCode{},
			Link:         "/miro/add-access-token",
			CodeExpected: 400,
		},
	}
	router := gin.Default()
	routes.ApplyRoutes(router)
	w := httptest.NewRecorder()

	for _, test := range tests {

		b, _ := json.Marshal(test.Oauth)
		req, _ := http.NewRequest("POST", test.Link, bytes.NewBuffer(b))
		router.ServeHTTP(w, req)
		assert.Equal(t, test.CodeExpected, w.Code)
	}
}

func TestErrorPostArea(t *testing.T) {
	router := gin.Default()
	routes.ApplyRoutes(router)
	w := httptest.NewRecorder()

	b, _ := json.Marshal(models.Comment{})
	req, _ := http.NewRequest("POST", "/areas", bytes.NewBuffer(b))
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestErrorGetArea(t *testing.T) {
	router := gin.Default()
	routes.ApplyRoutes(router)
	w := httptest.NewRecorder()

	b, _ := json.Marshal(models.Comment{})
	req, _ := http.NewRequest("GET", "/areas", bytes.NewBuffer(b))
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestErrorDeleteArea(t *testing.T) {
	router := gin.Default()
	routes.ApplyRoutes(router)
	w := httptest.NewRecorder()

	b, _ := json.Marshal(models.Comment{})
	req, _ := http.NewRequest("DELETE", "/areas", bytes.NewBuffer(b))
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}
