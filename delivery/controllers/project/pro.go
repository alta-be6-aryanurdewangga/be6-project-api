package project

import (
	"net/http"
	"part3/delivery/middlewares"
	"part3/lib/database/project"
	"part3/models/base"
	"part3/models/project/request"

	"github.com/labstack/echo/v4"
)

type ProController struct {
	repo project.Project
}

func New(repo project.Project) *ProController {
	return &ProController{
		repo: repo,
	}
}

func (pc *ProController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		user_id := int(middlewares.ExtractTokenId(c))
		newPro := request.ProRequest{}
		if err := c.Bind(&newPro); err != nil || newPro.Name_Pro == "" {
			return c.JSON(http.StatusBadRequest, base.BadRequest(nil, "error in input project", nil))
		}

		res, err := pc.repo.Create(user_id, newPro.ToProject())

		if err != nil {
			return c.JSON(http.StatusInternalServerError, base.InternalServerError(nil, "error in call database", nil))
		}

		return c.JSON(http.StatusCreated, base.Success(http.StatusCreated, "success create project", res.ToProResponse()))
	}
}
