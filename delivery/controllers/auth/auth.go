package auth

import (
	"net/http"
	"part3/delivery/middlewares"
	"part3/lib/database/auth"
	"part3/models/base"
	"part3/models/user/request"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	repo auth.Auth
}

func New(repo auth.Auth) *AuthController {
	return &AuthController{
		repo: repo,
	}
}

func (ac *AuthController) Login() echo.HandlerFunc{
	return func(c echo.Context) error {
		Userlogin := request.Userlogin{}

		if err := c.Bind(&Userlogin) ; err != nil {
			return c.JSON(http.StatusBadRequest,base.BadRequest(nil, "error in input file", nil))
		}
		
		checkedUser, err := ac.repo.Login(Userlogin)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, base.InternalServerError(nil, "error in call database", nil))
		}

		token, err := middlewares.GenerateToken(checkedUser)

		if err != nil {
			return c.JSON(http.StatusNotAcceptable, base.BadRequest(http.StatusNotAcceptable, "error in process token", nil))
		}

		return c.JSON(http.StatusOK, base.Success(nil, "success login", map[string]interface{}{
			"data" : checkedUser,
			"token":token,
		}))
	}
}