package task

import (
	"net/http"
	"part3/lib/database/task"
	"part3/models/base"
	"part3/models/task/request"

	"github.com/labstack/echo/v4"
)

type TaskController struct {
	repo task.Task
}

func New(repository task.Task) *TaskController {
	return &TaskController{
		repo: repository,
	}
}

func (tc *TaskController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		newTask := request.TaskRequest{}

		if err := c.Bind(&newTask); err != nil {
			return c.JSON(http.StatusBadRequest, base.BadRequest(
				http.StatusBadRequest,
				"error to create task",
				nil,
			))
		}

		res, err := tc.repo.Create(newTask.ToTask())

		if err != nil {
			return c.JSON(http.StatusInternalServerError, base.InternalServerError(
				http.StatusInternalServerError,
				"error in server process",
				nil,
			))
		}

		return c.JSON(http.StatusCreated, base.Success(
			http.StatusCreated,
			"success to create task",
			res,
		))
	}
}
