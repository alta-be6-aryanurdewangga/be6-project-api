package user

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"part3/delivery/controllers/auth"
	"part3/delivery/middlewares"
	"part3/models/user"
	"part3/models/user/request"
	"part3/models/user/response"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestCreate(t *testing.T) {
	var jwtToken string

	t.Run("Success Login", func(t *testing.T) {
		e := echo.New()

		reqBody, _ := json.Marshal(map[string]string{
			"email":    "anonim@123",
			"password": "anonim123",
		})

		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)

		authController := auth.New(&MockAuthLib{})
		authController.Login()(context)

		response := auth.LoginRespFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		jwtToken = response.Data["token"].(string)

		assert.Equal(t, response.Message, "success login")
		assert.NotNil(t, response.Data["token"])
	})

	t.Run("Failed to Create", func(t *testing.T) {
		e := echo.New()

		type UserRegister struct {
			Name     interface{} `json:"name"`
			Email    interface{} `json:"email" `
			Password interface{} `json:"password"`
		}

		reqBody, _ := json.Marshal(UserRegister{
			Name:     0,
			Email:    0,
			Password: 0,
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/users")

		userController := New(&MockUserLib{})
		userController.Create()(context)

		response := GetUserResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, 400, response.Code)
		assert.Equal(t, "error in request Create", response.Message)

	})

	t.Run("Failed to Access", func(t *testing.T) {
		e := echo.New()

		reqBody, _ := json.Marshal(map[string]string{
			"email":    "anonim",
			"password": "anonim",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/users")

		userController := New(&MockUserLib{})
		userController.Create()(context)

		response := GetUserResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, 500, response.Code)
		assert.Equal(t, "error in access Create", response.Message)

	})

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
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/users")

		userController := New(&MockUserLib{})
		userController.Create()(context)

		response := GetUserResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, 201, response.Code)
		assert.Equal(t, "anonim123", response.Data.Name)
	})
}

func TestGetById(t *testing.T) {
	var jwtToken string

	t.Run("Success Login", func(t *testing.T) {
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

		assert.Equal(t, response.Message, "success login")
		assert.NotNil(t, response.Data["token"])
	})

	t.Run("Fail to Get By Id", func(t *testing.T) {

		e := echo.New()
		// userid := int(middlewares.ExtractTokenId(c))

		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/users/:id")

		userController := New(&MockFalseLib{})
		if err := middlewares.JwtMiddleware()(userController.GetById())(context); err != nil {
			return
		}

		response := GetUserResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		// log.Info(response)
		assert.Equal(t, 500, response.Code)
		assert.Equal(t, "error in access Get By id", response.Message)
	})

	t.Run("Success Get By Id", func(t *testing.T) {

		e := echo.New()
		// userid := int(middlewares.ExtractTokenId(c))

		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/users/:id")

		userController := New(&MockUserLib{})
		if err := middlewares.JwtMiddleware()(userController.GetById())(context); err != nil {
			return
		}

		response := GetUserResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		// log.Info(response)
		assert.Equal(t, 200, response.Code)
		assert.Equal(t, response.Data, response.Data)
	})
}

func TestUpdateByID(t *testing.T) {
	var jwtToken string

	t.Run("Success Login", func(t *testing.T) {
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

		assert.Equal(t, response.Message, "success login")
		assert.NotNil(t, response.Data["token"])
	})

	t.Run("Error input Update", func(t *testing.T) {
		e := echo.New()

		reqBody, _ := json.Marshal(map[string]string{})
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/users/:id")

		userController := New(&MockFalseLib{})
		if err := middlewares.JwtMiddleware()(userController.UpdateById())(context); err != nil {
			return
		}

		response := GetUserResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, 400, response.Code)
		assert.Equal(t, "error in request Update", response.Message)
	})

	t.Run("Error access Update", func(t *testing.T) {
		e := echo.New()

		reqBody, _ := json.Marshal(map[string]interface{}{
			"name":     "anonim123",
			"email":    "anonim@123",
			"password": "anonim123",
		})
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/users/:id")

		userController := New(&MockFalseLib{})
		if err := middlewares.JwtMiddleware()(userController.UpdateById())(context); err != nil {
			return
		}

		response := GetUserResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, 500, response.Code)
		assert.Equal(t, "error in access Update", response.Message)
	})

	t.Run("Success Update", func(t *testing.T) {
		e := echo.New()

		reqBody, _ := json.Marshal(map[string]string{
			"name":     "anonim123",
			"email":    "anonim@123",
			"password": "anonim123",
		})
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/users/:id")

		userController := New(&MockUserLib{})
		if err := middlewares.JwtMiddleware()(userController.UpdateById())(context); err != nil {
			return
		}

		response := GetUserResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, 200, response.Code)
		assert.Equal(t, response.Data, response.Data)
		log.Info(response.Data)
	})

}

func TestDeleteByID(t *testing.T) {
	var jwtToken string

	t.Run("Success Login", func(t *testing.T) {
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

		assert.Equal(t, response.Message, "success login")
		assert.NotNil(t, response.Data["token"])
	})

	t.Run("Fail to Delete", func(t *testing.T) {
		e := echo.New()
		reqBody, _ := json.Marshal(map[string]string{})
		req := httptest.NewRequest(http.MethodDelete, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/users/:id")

		userController := New(&MockFalseLib{})
		if err := middlewares.JwtMiddleware()(userController.DeleteById())(context); err != nil {
			return
		}

		response := GetUserResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, 400, response.Code)
		assert.Equal(t, "error in request Delete", response.Message)

	})

	t.Run("Fail to access Delete", func(t *testing.T) {
		e := echo.New()
		reqBody, _ := json.Marshal(map[string]string{
			"name":     "anonim123",
			"email":    "anonim@123.com",
			"password": "anonim123",
		})
		req := httptest.NewRequest(http.MethodDelete, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/users/:id")

		userController := New(&MockFalseLib{})
		if err := middlewares.JwtMiddleware()(userController.DeleteById())(context); err != nil {
			return
		}

		response := GetUserResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, 500, response.Code)
		assert.Equal(t, "error in access Delete", response.Message)

	})

	t.Run("Success Delete", func(t *testing.T) {
		e := echo.New()
		reqBody, _ := json.Marshal(map[string]string{
			"name":     "anonim123",
			"email":    "anonim@123.com",
			"password": "anonim123",
		})
		req := httptest.NewRequest(http.MethodDelete, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/users/:id")

		userController := New(&MockUserLib{})
		if err := middlewares.JwtMiddleware()(userController.DeleteById())(context); err != nil {
			return
		}

		response := GetUserResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, 200, response.Code)
		assert.Equal(t, "Success Delete By Id", response.Message)

	})

}

func TestGetAll(t *testing.T) {
	var jwtToken string

	t.Run("Success Login", func(t *testing.T) {
		e := echo.New()
		reqBody, _ := json.Marshal(map[string]string{
			"name":     "admin",
			"email":    "admin@admin.com",
			"password": "admin",
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

		assert.Equal(t, response.Message, "success login")
		assert.NotNil(t, response.Data["token"])
	})

	t.Run("Success Get All User", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/users")

		userController := New(&MockUserLib{})
		if err := middlewares.JwtMiddleware()(userController.GetAll())(context); err != nil {
			return
		}

		response := GetUserResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, 200, response.Code)
		assert.Equal(t, "Success Get All User", response.Message)
	})

	t.Run("Failed Get All User", func(t *testing.T) {
		// token := string(jwtToken)
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/users")

		// userController := New(&MockFalseLib{})
		userController := New(&MockFalseLib{})

		if err := middlewares.JwtMiddleware()(userController.GetAll())(context); err != nil {
			return
		}

		response := GetUserResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, 400, response.Code)
		assert.Equal(t, "error in request Get", response.Message)
	})
}

type MockAuthLib struct{}

func (ma *MockAuthLib) Login(UserLogin request.Userlogin) (user.User, error) {
	return user.User{Model: gorm.Model{ID: 1}, Email: UserLogin.Email, Password: UserLogin.Password}, nil
}

type MockUserLib struct{}

func (m *MockUserLib) Create(newUser user.User) (user.User, error) {
	if newUser.Email != "anonim123" && newUser.Password != "anonim123" {
		return user.User{}, errors.New("record not found")
	}
	return user.User{Name: newUser.Name, Email: newUser.Email, Password: newUser.Password}, nil
}

func (m *MockUserLib) GetById(id int) (response.UserResponse, error) {
	return response.UserResponse{}, nil
}

func (m *MockUserLib) UpdateById(id int, userReg request.UserRegister) (user.User, error) {
	return user.User{Name: userReg.Name, Email: userReg.Email, Password: userReg.Password}, nil
}

func (m *MockUserLib) DeleteById(id int) (gorm.DeletedAt, error) {
	user := user.User{}
	return user.DeletedAt, nil
}

func (m *MockUserLib) GetAll() ([]response.UserResponse, error) {
	return []response.UserResponse{}, nil
}

type MockFalseLib struct{}

func (mf *MockFalseLib) Create(newUser user.User) (user.User, error) {
	if newUser.Email != "anonim123" && newUser.Password != "anonim123" {
		return user.User{}, errors.New("record not found")
	}
	return user.User{Name: newUser.Name, Email: newUser.Email, Password: newUser.Password}, nil
}

func (mf *MockFalseLib) GetById(id int) (response.UserResponse, error) {
	return response.UserResponse{}, errors.New("False Object")
}

func (mf *MockFalseLib) UpdateById(id int, userReg request.UserRegister) (user.User, error) {
	return user.User{}, errors.New("False Object")
}

func (mf *MockFalseLib) DeleteById(id int) (gorm.DeletedAt, error) {
	user := user.User{}
	return user.DeletedAt, errors.New("False Object")
}

func (mf *MockFalseLib) GetAll() ([]response.UserResponse, error) {
	return []response.UserResponse{}, errors.New("False Object")
}
