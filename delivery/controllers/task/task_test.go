package task

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"part3/delivery/controllers/auth"
	"part3/delivery/middlewares"
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
		assert.Equal(t, 200, response.Code)
		assert.NotNil(t, response.Data["token"])
	})

	t.Run("error in input task", func(t *testing.T) {
		e := echo.New()
		reqBody, _ := json.Marshal(map[string]string{
			"priority": "1",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/todo/tasks")

		taskController := New(&MockTaskLib{})
		// taskController.Create()(context)
		if err := middlewares.JwtMiddleware()(taskController.Create())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := GetTaskResponFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, 400, response.Code)
		assert.Equal(t, "error in input task", response.Message)
	})

	t.Run("error in database process", func(t *testing.T) {
		e := echo.New()
		reqBody, _ := json.Marshal(map[string]interface{}{
			"name_task": "anonim",
			"priority":  1,
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		log.Info(req)
		req.Header.Set("Content-Type", "application/json")
		// req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/todo/tasks")

		taskController := New(&MockTaskLib{})
		taskController.Create()(context)
		// if err := middlewares.JwtMiddleware()(taskController.Create())(context); err != nil {
		// 	log.Fatal(err)
		// 	return
		// }
		response := GetTaskResponFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, 500, response.Code)
		assert.Equal(t, "error in database process", response.Message)
	})
}

type MockTaskLib struct{}

func (m *MockTaskLib) Create(user_id int, newTask task.Task) (task.Task, error) {
	if newTask.Name_Task != "anonim123" {
		return task.Task{}, errors.New("error in database process")
	}

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
