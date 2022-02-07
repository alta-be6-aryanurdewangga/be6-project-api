package auth

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"part3/models/user"
	"part3/models/user/request"
	"testing"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func TestLogin(t *testing.T) {
	t.Run("erro in input file", func(t *testing.T) {
		e := echo.New()

		reqBody, _ := json.Marshal(map[string]string{
			"name" : "anonim123",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/login")


	})
}

type MockAuthLib struct{}

func (m *MockAuthLib) Login(UserLogin request.Userlogin) (user.User, error) {
	return user.User{Model: gorm.Model{ID: 1}, Name: "anonim123", Email: "anonim@123", Password: "anonim123"}, nil
}
