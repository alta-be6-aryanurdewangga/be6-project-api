package task

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"part3/delivery/controllers/auth"
	"part3/models/task"
	"part3/models/user"
	"part3/models/user/request"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestCreate(t *testing.T) {
	var jwtToken string
	t.Run("success login", func(t *testing.T) {
		e := echo.New()
		reqBody, _ := json.Marshal(map[string]string{
			"email":    "anonim@123",
			"password": "anonim123",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/login")
		authController := auth.New(&MockAuthLib{})
		authController.Login()(context)
		response := auth.LoginRespFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		jwtToken = response.Data["token"].(string)
		log.Info(jwtToken)
		assert.Equal(t, 200, response.Code)
		assert.NotNil(t, response.Data["token"])
	})
}

type MockTaskLib struct{}

func (m *MockTaskLib) Create(user_id int, newTask task.Task) (task.Task, error) {
	return task.Task{
		User_ID:   uint(user_id),
		Name_Task: newTask.Name_Task,
		Priority:  newTask.Priority,
	}, nil
}

type MockAuthLib struct{}

func (m *MockAuthLib) Login(UserLogin request.Userlogin) (user.User, error) {
	if UserLogin.Email != "anonim@123" && UserLogin.Password != "anonim123" {
		return user.User{}, errors.New("record not found")
	}
	return user.User{Model: gorm.Model{ID: 1}, Email: UserLogin.Email, Password: UserLogin.Password}, nil
}
