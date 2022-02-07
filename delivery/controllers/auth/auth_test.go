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
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

func TestLogin(t *testing.T) {
	t.Run("error in input file", func(t *testing.T) {
		e := echo.New()

		reqBody, _ := json.Marshal(map[string]string{
			"name" : "anonim123",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		log.Info(req)
		res := httptest.NewRecorder()
		log.Info(res)
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/login")
		log.Info(context)
		authCont := New(&MockAuthLib{})
		authCont.Login()(context)
		log.Info(authCont)
		resp := LoginRespFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &resp)

		log.Info(resp)

		log.Info(resp.Data.Name)
	})
}

type MockAuthLib struct{}

func (m *MockAuthLib) Login(UserLogin request.Userlogin) (user.User, error) {
	return user.User{Model: gorm.Model{ID: 1}, Email: UserLogin.Email, Password: UserLogin.Password}, nil
}
