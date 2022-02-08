package task

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"part3/models/task"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	t.Run("Success Create", func(t *testing.T) {
		e := echo.New()
		reqBody, _ := json.Marshal(map[string]interface{}{
			"name_task": "task1",
			"priority":  1,
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/tasks")

		taskController := New(&MockTaskLib{})
		taskController.Create()(context)

		resp := GetTaskResponFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &resp)
		assert.Equal(t, resp.Name_Task, resp.Name_Task)
		assert.Equal(t, resp.P, resp.Name_Task)
	})
}

type MockTaskLib struct{}

func (m *MockTaskLib) Create(newTask task.Task) (task.Task, error) {
	return task.Task{
		Name_Task: newTask.Name_Task,
		Priority:  newTask.Priority,
	}, nil
}
