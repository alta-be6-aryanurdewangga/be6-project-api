package task

import (
	"net/http"
	"part3/delivery/middlewares"
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
		user_id := int(middlewares.ExtractTokenId(c))
		newTask := request.TaskRequest{}

		if err := c.Bind(&newTask); err != nil || newTask.Name_Task == "" {
			return c.JSON(http.StatusBadRequest, base.BadRequest(
				http.StatusBadRequest,
				"error in input task",
				nil,
			))
		}

		res, err := tc.repo.Create(user_id,newTask.ToTask())

		if err != nil {
			return c.JSON(http.StatusInternalServerError, base.InternalServerError(
				http.StatusInternalServerError,
				"error in database process",
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
