package project

import (
	"net/http"
	"part3/delivery/middlewares"
	"part3/lib/database/project"
	"part3/models/base"
	"part3/models/project/request"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProController struct {
	repo project.Project
}

func NewRepo(repo project.Project) *ProController {
	return &ProController{
		repo: repo,
	}
}

func (pc *ProController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		user_id := int(middlewares.ExtractTokenId(c))
		newPro := request.ProRequest{}
		if err := c.Bind(&newPro); err != nil || newPro.Name == "" {
			return c.JSON(http.StatusBadRequest, base.BadRequest(nil, "error in input project", nil))
		}

		res, err := pc.repo.Create(user_id, newPro.ToProject())

		if err != nil {
			return c.JSON(http.StatusInternalServerError, base.InternalServerError(nil, "error in call database", nil))
		}
		return c.JSON(http.StatusCreated, base.Success(http.StatusCreated, "success create project", res.ToProResponse()))
	}
}

func (pc *ProController) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		user_id := int(middlewares.ExtractTokenId(c))

		res, err := pc.repo.GetAll(user_id)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, base.InternalServerError(
				http.StatusInternalServerError,
				"error in database process",
				nil,
			))
		}

		return c.JSON(http.StatusCreated, base.Success(
			http.StatusCreated,
			"success to get all project",
			res,
		))
	}
}

func (pc *ProController) Put() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		user_id := int(middlewares.ExtractTokenId(c))
		upPro := request.ProRequest{}
		if err := c.Bind(&upPro); err != nil || upPro.Name == "" {
			return c.JSON(http.StatusBadRequest, base.BadRequest(nil, "error in input project", nil))
		}

		res, err := pc.repo.UpdateById(id, user_id, upPro)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, base.InternalServerError(
				http.StatusInternalServerError,
				"error in database proces",
				nil,
			))
		}

		return c.JSON(http.StatusCreated, base.Success(
			http.StatusCreated,
			"success to update project",
			res.ToProResponse(),
		))
	}
}

func (pc *ProController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))

		user_id := int(middlewares.ExtractTokenId(c))

		res, err := pc.repo.DeleteById(id, user_id)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, base.InternalServerError(
				http.StatusInternalServerError,
				"error in database process",
				nil,
			))
		}

		return c.JSON(http.StatusOK, base.Success(
			nil,
			"success to delete project",
			res,
		))
	}
}
