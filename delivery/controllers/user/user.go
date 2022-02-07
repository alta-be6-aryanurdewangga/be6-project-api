package user

import (
	"net/http"
	"part3/lib/database/user"
	"part3/models/base"
	"part3/models/user/request"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	repo user.User
}

func New(repository user.User) *UserController {
	return &UserController{
		repo: repository,
	}
}

func (uc *UserController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		newUser := request.UserRegister{}

		if err := c.Bind(&newUser); err != nil {
			return c.JSON(http.StatusBadRequest, base.BadRequest(nil, "error in request Create", nil))
		}

		res, err := uc.repo.Create(newUser.ToUser())

		if err != nil {
			return c.JSON(http.StatusInternalServerError, base.InternalServerError(nil, "error in access Create", nil))
		}

		return c.JSON(http.StatusCreated, base.Success(http.StatusCreated, "Success Create", res))
	}
}