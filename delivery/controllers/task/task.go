package task

import (
	"net/http"
	"part3/delivery/middlewares"
	"part3/lib/database/task"
	"part3/models/base"
	t "part3/models/task"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
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
		newTask := TaskRequest{}
		log.Info(c)
		if err := c.Bind(&newTask); err != nil /* || newTask.Name_Task == ""  */ {
			log.Info(err)
			return c.JSON(http.StatusBadRequest, base.BadRequest(
				http.StatusBadRequest,
				"error in input task",
				nil,
			))
		}

		res, err := tc.repo.Create(user_id, t.Task{Name_Task: newTask.Name_Task, Priority: newTask.Priority})

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
