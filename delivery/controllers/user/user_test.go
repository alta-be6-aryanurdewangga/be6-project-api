package user

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"part3/models/user"

	// lib "part3/lib/database/user"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	t.Run("Success Create", func(t *testing.T) {
		e := echo.New()
		reqBody, _ := json.Marshal(map[string]string{
			"name":     "anonim123",
			"email":    "anonim@123",
			"password": "anonim123",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/users")

		userController := New(&MockUserLib{})
		userController.Create()(context)

		response := GetUserResponseFormat{}
		
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		log.Info(response)
		assert.Equal(t, 201, response.Code)
		assert.Equal(t, "anonim123", response.Data.Name)
	})
}

type MockUserLib struct{}

func (m *MockUserLib) Create(newUser user.User) (user.User, error) {
	return user.User{Name: newUser.Name, Email: newUser.Email, Password: newUser.Password}, nil
}
