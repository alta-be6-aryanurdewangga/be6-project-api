package project

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"part3/delivery/controllers/auth"
	"part3/delivery/middlewares"
	proMod "part3/models/project"
	"part3/models/project/request"
	"part3/models/project/response"
	"part3/models/user"
	reqU "part3/models/user/request"
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

	t.Run("error in input project", func(t *testing.T) {
		e := echo.New()
		reqBody, _ := json.Marshal(map[string]int{"name_pro": 1})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content=Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/projects")

		taskController := NewRepo(&MockFailProLib{})
		if err := middlewares.JwtMiddleware()(taskController.Create())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := GetRespFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, 400, response.Code)
		assert.Equal(t, "error in input project", response.Message)
	})

	t.Run("error in database process", func(t *testing.T) {
		e := echo.New()
		reqBody, _ := json.Marshal(map[string]interface{}{
			"name_pro": "anonim",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/projects")

		taskController := NewRepo(&MockFailProLib{})
		// taskController.Create()(context)
		if err := middlewares.JwtMiddleware()(taskController.Create())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := GetRespFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, 500, response.Code)
		assert.Equal(t, "error in call database", response.Message)
	})

	t.Run("success to create project", func(t *testing.T) {
		e := echo.New()
		reqBody, _ := json.Marshal(map[string]interface{}{
			"name_pro": "anonim",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/projects")

		taskController := NewRepo(&MockProLib{})
		// taskController.Create()(context)
		if err := middlewares.JwtMiddleware()(taskController.Create())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := GetRespFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, 201, response.Code)
		assert.Equal(t, "success create project", response.Message)
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
		context.SetPath("/projects/")

		taskController := NewRepo(&MockFailProLib{})

		if err := middlewares.JwtMiddleware()(taskController.GetAll())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := GetRespFormat{}
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
		context.SetPath("/projects/")

		taskController := NewRepo(&MockProLib{})

		if err := middlewares.JwtMiddleware()(taskController.GetAll())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := GetRespFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, 201, response.Code)
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

	t.Run("error in input projects", func(t *testing.T) {
		e := echo.New()
		reqBody, _ := json.Marshal(map[string]interface{}{})
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/projects/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")
		taskController := NewRepo(&MockFailProLib{})
		// taskController.Create()(context)
		if err := middlewares.JwtMiddleware()(taskController.Put())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := GetRespFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, 400, response.Code)
		assert.Equal(t, "error in input project", response.Message)
	})

	t.Run("success to update project", func(t *testing.T) {
		e := echo.New()
		reqBody, _ := json.Marshal(map[string]interface{}{
			"name_pro": "anonim123",
		})
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/todo/tasks/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")
		log.Info(context.Path())
		ProkController := NewRepo(&MockProLib{})
		// taskController.Create()(context)
		if err := middlewares.JwtMiddleware()(ProkController.Put())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := GetRespFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, 201, response.Code)
		assert.Equal(t, "success to update project", response.Message)
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
		req := httptest.NewRequest(http.MethodDelete, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/todo/tasks/1")
		taskController := NewRepo(&MockFailProLib{})
		// taskController.Create()(context)
		if err := middlewares.JwtMiddleware()(taskController.Delete())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := GetRespFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, 500, response.Code)
		assert.Equal(t, "error in database process", response.Message)
	})

	t.Run("success to delete project", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodDelete, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/todo/tasks/1")

		taskController := NewRepo(&MockProLib{})
		// taskController.Create()(context)
		if err := middlewares.JwtMiddleware()(taskController.Delete())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := GetRespFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, 200, response.Code)
		assert.Equal(t, "success to delete project", response.Message)
	})
}

type MockAuthLib struct{}

func (m *MockAuthLib) Login(UserLogin reqU.Userlogin) (user.User, error) {
	return user.User{Model: gorm.Model{ID: 1}, Email: UserLogin.Email, Password: UserLogin.Password}, nil
}

type MockProLib struct{}

func (m *MockProLib) Create(user_id int, newPro proMod.Project) (proMod.Project, error) {
	return proMod.Project{User_ID: uint(user_id), Name_Pro: newPro.Name_Pro}, nil
}

func (m *MockProLib) GetAll(user_id int) ([]response.ProResponse, error) {
	return []response.ProResponse{}, nil
}

func (m *MockProLib) UpdateById(id int, user_id int, upPro request.ProRequest) (proMod.Project, error) {
	return proMod.Project{User_ID: uint(user_id), Name_Pro: upPro.Name_Pro}, nil
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

func (m *MockFailProLib) GetAll(user_id int) ([]response.ProResponse, error) {
	return []response.ProResponse{}, errors.New("error in call database")
}

func (m *MockFailProLib) UpdateById(id int, user_id int, upPro request.ProRequest) (proMod.Project, error) {
	return proMod.Project{User_ID: uint(user_id), Name_Pro: upPro.Name_Pro}, errors.New("error in call database")
}

func (m *MockFailProLib) DeleteById(id int, user_id int) (gorm.DeletedAt, error) {
	pro := proMod.Project{}
	return pro.DeletedAt, errors.New("error in database process")
}

func (m *MockFailProLib) GetById(id int, user_id int) (proMod.Project, error) {
	return proMod.Project{}, nil
}
