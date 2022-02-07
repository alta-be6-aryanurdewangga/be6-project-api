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
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestLogin(t *testing.T) {
	t.Run("error in input file", func(t *testing.T) {
		e := echo.New()

		reqBody, _ := json.Marshal(map[string]string{
			"email": "anonim123",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/login")
		authCont := New(&MockAuthLib{})
		authCont.Login()(context)
		resp := LoginRespFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		assert.Equal(t, 400, resp.Code)
		assert.Equal(t, "error in input file", resp.Message)
	})

	t.Run("error in call database", func(t *testing.T) {
		
	})
}

type MockAuthLib struct{}

func (m *MockAuthLib) Login(UserLogin request.Userlogin) (user.User, error) {
	return user.User{Model: gorm.Model{ID: 1}, Email: UserLogin.Email, Password: UserLogin.Password}, nil
}
