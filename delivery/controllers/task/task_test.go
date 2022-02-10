package task

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"part3/delivery/controllers/auth"
	"part3/delivery/middlewares"
	proMod "part3/models/project"
	proReq "part3/models/project/request"
	proResp "part3/models/project/response"
	"part3/models/task"
	"part3/models/task/request"
	"part3/models/task/response"
	"part3/models/user"
	reqU "part3/models/user/request"
	"testing"

	"github.com/labstack/echo/v4"
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
		reqBody, _ := json.Marshal(map[string]interface{}{})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/todo/tasks")

		taskController := New(&MockFailTaskLib{}, &MockFailProLib{})
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
			"name":       "anonim",
			"priority":   1,
			"project_id": 1,
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/todo/tasks")

		taskController := New(&MockFailTaskLib{}, &MockFailProLib{})
		// taskController.Create()(context)
		if err := middlewares.JwtMiddleware()(taskController.Create())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := GetTaskResponFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, 500, response.Code)
		assert.Equal(t, "error in database process", response.Message)
	})

	t.Run("error in database process", func(t *testing.T) {
		e := echo.New()
		reqBody, _ := json.Marshal(map[string]interface{}{
			"name":       "anonim",
			"priority":   1,
			"project_id": 1,
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/todo/tasks")

		taskController := New(&MockFailTaskLib{}, &MockProLib{})
		// taskController.Create()(context)
		if err := middlewares.JwtMiddleware()(taskController.Create())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := GetTaskResponFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, 500, response.Code)
		assert.Equal(t, "error in database process", response.Message)
	})

	t.Run("success to create task", func(t *testing.T) {
		e := echo.New()
		reqBody, _ := json.Marshal(map[string]interface{}{
			"name":       "anonim",
			"priority":   1,
			"project_id": 1,
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/todo/tasks")
		taskController := New(&MockTaskLib{}, &MockProLib{})
		// taskController.Create()(context)
		if err := middlewares.JwtMiddleware()(taskController.Create())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := GetTaskResponFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, 201, response.Code)
		assert.Equal(t, "success to create task", response.Message)
	})
}

func TestGetAll(t *testing.T) {
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

	t.Run("error in database process", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/todo/tasks")

		taskController := New(&MockFailTaskLib{}, &MockProLib{})

		if err := middlewares.JwtMiddleware()(taskController.GetAll())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := GetTaskResponFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, 500, response.Code)
		assert.Equal(t, "error in database process", response.Message)
	})

	t.Run("success to get all task", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/todo/tasks")

		taskController := New(&MockTaskLib{}, &MockProLib{})

		if err := middlewares.JwtMiddleware()(taskController.GetAll())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := GetTaskResponFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, 200, response.Code)
		assert.Equal(t, "success to get all task", response.Message)
	})
}

func TestPut(t *testing.T) {
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
		reqBody, _ := json.Marshal(map[string]interface{}{})
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/todo/tasks/1")

		taskController := New(&MockFailTaskLib{}, &MockFailProLib{})
		// taskController.Create()(context)
		if err := middlewares.JwtMiddleware()(taskController.Put())(context); err != nil {
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
			"name": "anonim",
			"priority":  1,
			"project_id":5,
		})
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/todo/tasks/1")
		taskController := New(&MockFailTaskLib{}, &MockFailProLib{})
		// taskController.Create()(context)
		if err := middlewares.JwtMiddleware()(taskController.Put())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := GetTaskResponFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, 500, response.Code)
		assert.Equal(t, "error in database process", response.Message)
	})

	t.Run("error in database process", func(t *testing.T) {
		e := echo.New()
		reqBody, _ := json.Marshal(map[string]interface{}{
			"name": "anonim",
			"priority":  1,
			"project_id":5,
		})
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/todo/tasks/1")
		taskController := New(&MockFailTaskLib{}, &MockProLib{})
		// taskController.Create()(context)
		if err := middlewares.JwtMiddleware()(taskController.Put())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := GetTaskResponFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, 500, response.Code)
		assert.Equal(t, "error in database process", response.Message)
	})

	t.Run("success to update task", func(t *testing.T) {
		e := echo.New()
		reqBody, _ := json.Marshal(map[string]interface{}{
			"name": "anonim",
			"priority":  1,
			"project_id":5,
		})
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/todo/tasks/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")

		taskController := New(&MockTaskLib{}, &MockProLib{})
		// taskController.Create()(context)
		if err := middlewares.JwtMiddleware()(taskController.Put())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := GetTaskResponFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, 200, response.Code)
		assert.Equal(t, "success to update task", response.Message)
	})
}

func TestDelete(t *testing.T) {
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

	t.Run("error in database process", func(t *testing.T) {
		e := echo.New()
		reqBody, _ := json.Marshal(map[string]interface{}{
			"name_task": "anonim",
			"priority":  1,
		})
		req := httptest.NewRequest(http.MethodDelete, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/todo/tasks/1")
		taskController := New(&MockFailTaskLib{}, &MockProLib{})
		// taskController.Create()(context)
		if err := middlewares.JwtMiddleware()(taskController.Delete())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := GetTaskResponFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, 500, response.Code)
		assert.Equal(t, "error in database process", response.Message)
	})

	t.Run("success to delete task", func(t *testing.T) {
		e := echo.New()
		reqBody, _ := json.Marshal(map[string]interface{}{
			"name_task": "anonim123",
			"priority":  1,
		})
		req := httptest.NewRequest(http.MethodDelete, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/todo/tasks/1")

		taskController := New(&MockTaskLib{}, &MockProLib{})
		// taskController.Create()(context)
		if err := middlewares.JwtMiddleware()(taskController.Delete())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := GetTaskResponFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, 200, response.Code)
		assert.Equal(t, "success to delete task", response.Message)
	})
}

type MockTaskLib struct{}

func (m *MockTaskLib) Create(user_id int, newTask task.Task) (task.Task, error) {

	return task.Task{
		User_ID:  uint(user_id),
		Name:     newTask.Name,
		Priority: newTask.Priority,
	}, nil
}

func (m *MockTaskLib) GetAll(user_id int) ([]response.TaskResponse, error) {

	return []response.TaskResponse{}, nil
}

func (m *MockTaskLib) UpdateById(id int, user_id int, taskReg request.TaskRequest) (task.Task, error) {

	return task.Task{}, nil
}

func (m *MockTaskLib) DeleteById(id int, user_id int) (gorm.DeletedAt, error) {
	task := task.Task{}
	return task.DeletedAt, nil
}

type MockFailTaskLib struct{}

func (mf *MockFailTaskLib) Create(user_id int, newTask task.Task) (task.Task, error) {

	return task.Task{}, errors.New("error in database process")
}

func (mf *MockFailTaskLib) GetAll(user_id int) ([]response.TaskResponse, error) {
	return []response.TaskResponse{}, errors.New("error in database process")
}

func (mf *MockFailTaskLib) UpdateById(id int, user_id int, taskReg request.TaskRequest) (task.Task, error) {

	return task.Task{}, errors.New("error in database process")
}

func (m *MockFailTaskLib) DeleteById(id int, user_id int) (gorm.DeletedAt, error) {
	task := task.Task{}
	return task.DeletedAt, errors.New("error in database process")
}

/* Moch authentification */
type MockAuthLib struct{}

func (m *MockAuthLib) Login(UserLogin reqU.Userlogin) (user.User, error) {
	if UserLogin.Email != "anonim@123" && UserLogin.Password != "anonim123" {
		return user.User{}, errors.New("record not found")
	}
	return user.User{Model: gorm.Model{ID: 1}, Email: UserLogin.Email, Password: UserLogin.Password}, nil
}

type MockProLib struct{}

func (m *MockProLib) Create(user_id int, newPro proMod.Project) (proMod.Project, error) {
	return proMod.Project{User_ID: uint(user_id), Name: newPro.Name}, nil
}

func (m *MockProLib) GetAll(user_id int) ([]proResp.ProResponse, error) {
	return []proResp.ProResponse{}, nil
}

func (m *MockProLib) UpdateById(id int, user_id int, upPro proReq.ProRequest) (proMod.Project, error) {
	return proMod.Project{User_ID: uint(user_id), Name: upPro.Name}, nil
}

func (m *MockProLib) DeleteById(id int, user_id int) (gorm.DeletedAt, error) {
	pro := proMod.Project{}
	return pro.DeletedAt, nil
}

func (m *MockProLib) GetById(id int, user_id int) (proMod.Project, error) {
	return proMod.Project{}, nil
}

type MockFailProLib struct{}

func (m *MockFailProLib) Create(user_id int, newPro proMod.Project) (proMod.Project, error) {
	return proMod.Project{}, errors.New("error in call database")
}

func (m *MockFailProLib) GetAll(user_id int) ([]proResp.ProResponse, error) {
	return []proResp.ProResponse{}, errors.New("error in call database")
}

func (m *MockFailProLib) UpdateById(id int, user_id int, upPro proReq.ProRequest) (proMod.Project, error) {
	return proMod.Project{User_ID: uint(user_id), Name: upPro.Name}, errors.New("error in call database")
}

func (m *MockFailProLib) DeleteById(id int, user_id int) (gorm.DeletedAt, error) {
	pro := proMod.Project{}
	return pro.DeletedAt, errors.New("error in database process")
}

func (m *MockFailProLib) GetById(id int, user_id int) (proMod.Project, error) {
	return proMod.Project{},  errors.New("error in database process")
}
