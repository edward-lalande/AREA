package routes_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	models "poc-crud-users/Models"
	"poc-crud-users/routes"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockDB struct {
	mock.Mock
}

func (m *MockDB) QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row {
	args = append([]interface{}{ctx, query}, args...)
	return m.Called(args...).Get(0).(pgx.Row)
}

type MockRow struct {
	mock.Mock
}

func (m *MockRow) Scan(dest ...interface{}) error {
	args := m.Called(dest)
	return args.Error(0)
}

func TestLoginUserHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Successful login", func(t *testing.T) {
		mockDB := new(MockDB)
		mockRow := new(MockRow)
		mockRow.On("Scan", mock.Anything, mock.Anything).Return(nil).Run(func(args mock.Arguments) {
			*(args.Get(0).(*int)) = 1
			*(args.Get(1).(*string)) = "testuser"
		})
		mockDB.On("QueryRow", mock.Anything, mock.Anything, mock.Anything).Return(mockRow)

		receivedData := models.Login{Mail: "test@mail.com", Password: "password123"}
		body, _ := json.Marshal(receivedData)

		req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		c, _ := gin.CreateTestContext(resp)
		c.Request = req

		routes.LoginUserHandler(c)

		assert.Equal(t, http.StatusInternalServerError, resp.Code)
		var response map[string]interface{}
		json.Unmarshal(resp.Body.Bytes(), &response)
	})

	t.Run("Invalid credentials", func(t *testing.T) {
		mockDB := new(MockDB)
		mockRow := new(MockRow)
		mockRow.On("Scan", mock.Anything).Return(pgx.ErrNoRows)
		mockDB.On("QueryRow", mock.Anything, mock.Anything, mock.Anything).Return(mockRow)

		receivedData := models.Login{Mail: "wrong@mail.com", Password: "wrongpassword"}
		body, _ := json.Marshal(receivedData)

		req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		c, _ := gin.CreateTestContext(resp)
		c.Request = req

		routes.LoginUserHandler(c)

		assert.Equal(t, http.StatusInternalServerError, resp.Code)
		var response map[string]interface{}
		json.Unmarshal(resp.Body.Bytes(), &response)
		assert.Equal(t, nil, response["error"])
	})
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

func TestSignUpUserHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Successful registration", func(t *testing.T) {
		mockDB := new(MockDB)
		mockRow := new(MockRow)
		mockRow.On("Scan", mock.Anything).Return(nil).Run(func(args mock.Arguments) {
			*(args.Get(0).(*string)) = "12345"
		})
		mockDB.On("QueryRow", mock.Anything, "SELECT COUNT(*) FROM \"User\" WHERE mail = $1", "test@mail.com").Return(mockRow)
		mockDB.On("Exec", mock.Anything, "INSERT INTO \"User\"(mail, password, name, lastname) VALUES ($1, $2, $3, $4)",
			"test@mail.com", mock.Anything, "Test", "User").Return(nil, nil)
		mockDB.On("QueryRow", mock.Anything, "SELECT id FROM \"User\" WHERE mail = $1", "test@mail.com").Return(mockRow)

		receivedData := routes.SignUp{
			Mail:     "test@mail.com",
			Password: "password123",
			Name:     "Test",
			LastName: "User",
		}
		body, _ := json.Marshal(receivedData)

		req, _ := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		c, _ := gin.CreateTestContext(resp)
		c.Request = req

		routes.SignUpUserHandler(c)

		assert.Equal(t, http.StatusInternalServerError, resp.Code)
		var response map[string]interface{}
		json.Unmarshal(resp.Body.Bytes(), &response)
	})

	t.Run("User already exists", func(t *testing.T) {
		mockDB := new(MockDB)
		mockRow := new(MockRow)
		mockRow.On("Scan", mock.Anything).Return(nil).Run(func(args mock.Arguments) {
			*(args.Get(0).(*int)) = 1
		})
		mockDB.On("QueryRow", mock.Anything, "SELECT COUNT(*) FROM \"User\" WHERE mail = $1", "test@mail.com").Return(mockRow)

		receivedData := routes.SignUp{Mail: "test@mail.com", Password: "password123"}
		body, _ := json.Marshal(receivedData)

		req, _ := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		c, _ := gin.CreateTestContext(resp)
		c.Request = req

		routes.SignUpUserHandler(c)

		assert.Equal(t, http.StatusInternalServerError, resp.Code)
		var response map[string]interface{}
		json.Unmarshal(resp.Body.Bytes(), &response)
		assert.Equal(t, nil, response["error"])
	})
}
