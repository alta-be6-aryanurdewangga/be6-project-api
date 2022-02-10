package task

import (
	"net/http"
	"part3/delivery/middlewares"
	"part3/lib/database/project"
	"part3/lib/database/task"
	"part3/models/base"
	"part3/models/task/request"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type TaskController struct {
	repo    task.Task
	proLib  project.Project
}

func New(repository task.Task, proLib project.Project) *TaskController {
	return &TaskController{
		repo:    repository,
		proLib:  proLib,
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

		log.Info(tc.proLib)
		log.Info(newTask.Project_id)
		log.Info(tc.proLib.GetById(int(newTask.Project_id), user_id))
		if tc.proLib == nil {
			log.Fatal()
		}
		if _, err := tc.proLib.GetById(int(newTask.Project_id), user_id); err != nil {
			return c.JSON(http.StatusInternalServerError, base.InternalServerError(
				http.StatusInternalServerError,
				"error in database process",
				nil,
			))
		}

		res, err := tc.repo.Create(user_id, newTask.ToTask())

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
			res.ToTaskResponse(),
		))
	}
}

// func (tc *TaskController) GetAll() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		user_id := int(middlewares.ExtractTokenId(c))

// 		res, err := tc.repo.GetAll(user_id)

// 		if err != nil {
// 			return c.JSON(http.StatusInternalServerError, base.InternalServerError(
// 				http.StatusInternalServerError,
// 				"error in database process",
// 				nil,
// 			))
// 		}

// 		return c.JSON(http.StatusCreated, base.Success(
// 			http.StatusCreated,
// 			"success to get all task",
// 			res,
// 		))
// 	}
// }

// func (tc *TaskController) Put() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		id, _ := strconv.Atoi(c.Param("id"))
// 		user_id := int(middlewares.ExtractTokenId(c))
// 		upTask := request.TaskRequest{}
// 		if err := c.Bind(&upTask); err != nil || upTask.Name_Task == "" {
// 			return c.JSON(http.StatusBadRequest, base.BadRequest(
// 				http.StatusBadRequest,
// 				"error in input task",
// 				nil,
// 			))
// 		}

// 		res, err := tc.repo.UpdateById(id, user_id, upTask)

// 		if err != nil {
// 			return c.JSON(http.StatusInternalServerError, base.InternalServerError(
// 				http.StatusInternalServerError,
// 				"error in database process",
// 				nil,
// 			))
// 		}

// 		return c.JSON(http.StatusCreated, base.Success(
// 			http.StatusCreated,
// 			"success to update task",
// 			res.ToTaskResponse(),
// 		))
// 	}
// }

// func (tc *TaskController) Delete() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		id, _ := strconv.Atoi(c.Param("id"))

// 		user_id := int(middlewares.ExtractTokenId(c))
// 		upTask := request.TaskRequest{}
// 		if err := c.Bind(&upTask); err != nil || upTask.Name_Task == "" {
// 			return c.JSON(http.StatusBadRequest, base.BadRequest(
// 				http.StatusBadRequest,
// 				"error in input task",
// 				nil,
// 			))
// 		}

// 		res, err := tc.repo.DeleteById(id, user_id)

// 		if err != nil {
// 			return c.JSON(http.StatusInternalServerError, base.InternalServerError(
// 				http.StatusInternalServerError,
// 				"error in database process",
// 				nil,
// 			))
// 		}

// 		return c.JSON(http.StatusCreated, base.Success(
// 			http.StatusCreated,
// 			"success to delete task",
// 			res,
// 		))
// 	}
// }
