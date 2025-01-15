package oauth_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	models "google/Models"
	"google/oauth"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCallBack(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name         string
		queryParams  string
		expectedCode int
		expectedURL  string
	}{
		{
			name:         "Valid code",
			queryParams:  "?code=valid_code",
			expectedCode: http.StatusFound,
			expectedURL:  "http://localhost:8081/login?google_code=valid_code",
		},
		{
			name:         "Missing code",
			queryParams:  "",
			expectedCode: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := gin.New()
			router.GET("/callback", oauth.CallBack)
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/callback"+tt.queryParams, nil)

			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedCode, w.Code)

			if tt.expectedCode == http.StatusFound {
				assert.Equal(t, tt.expectedURL, w.Header().Get("Location"))
			}
		})
	}
}

func TestAddCallBack(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name         string
		queryParams  string
		expectedCode int
		expectedURL  string
	}{
		{
			name:         "Valid code",
			queryParams:  "?code=valid_code",
			expectedCode: http.StatusFound,
			expectedURL:  "http://localhost:8081/login?google_code=valid_code",
		},
		{
			name:         "Missing code",
			queryParams:  "",
			expectedCode: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := gin.New()
			router.GET("/add-callback", oauth.CallBack)
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/add-callback"+tt.queryParams, nil)

			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedCode, w.Code)

			if tt.expectedCode == http.StatusFound {
				assert.Equal(t, tt.expectedURL, w.Header().Get("Location"))
			}
		})
	}
}

func TestAddAccessToken(t *testing.T) {
	gin.SetMode(gin.TestMode)

	validToken := "valid_token"
	tests := []struct {
		name         string
		headerToken  string
		requestBody  models.OauthInformation
		expectedCode int
	}{
		{
			name:        "Valid request",
			headerToken: validToken,
			requestBody: models.OauthInformation{
				Code: "valid_code",
			},
			expectedCode: http.StatusInternalServerError,
		},
		{
			name:        "Missing token",
			headerToken: "",
			requestBody: models.OauthInformation{
				Code: "valid_code",
			},
			expectedCode: http.StatusInternalServerError,
		},
		{
			name:        "Invalid token",
			headerToken: "invalid_token",
			requestBody: models.OauthInformation{
				Code: "valid_code",
			},
			expectedCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := gin.New()
			router.POST("/access-token", oauth.AddAccessToken)
			w := httptest.NewRecorder()

			jsonData, _ := json.Marshal(tt.requestBody)
			req, _ := http.NewRequest("POST", "/access-token", bytes.NewBuffer(jsonData))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("token", tt.headerToken)
			router.ServeHTTP(w, req)
			assert.Equal(t, tt.expectedCode, w.Code)
		})
	}
}

func TestOAuthFront(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Get OAuth URL", func(t *testing.T) {
		router := gin.New()
		router.GET("/oauth", oauth.OAuthFront)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/oauth", nil)

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		assert.Contains(t, w.Body.String(), "https://accounts.google.com/o/oauth2/auth")
		assert.Contains(t, w.Body.String(), "response_type=code")
	})
}

func TestAddOAuthFront(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Get OAuth URL", func(t *testing.T) {
		router := gin.New()
		router.GET("/oauth", oauth.OAuthFront)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/oauth", nil)

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		assert.Contains(t, w.Body.String(), "https://accounts.google.com/o/oauth2/auth")
	})
}
