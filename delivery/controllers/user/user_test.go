package user

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"part3/models/user"
	lib "part3/lib/database/user"
	"testing"

	"github.com/labstack/echo/v4"

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

		response := 

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		
	})
}

type MockUserLib struct{}

func (m *MockUserLib) Create(newUser user.User) (user.User, error) {
	return user.User{Name: "anonim123", Email: "anonim@123", Password: "anonim123"}, nil
}
