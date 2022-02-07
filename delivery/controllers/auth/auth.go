package auth

import (
	"part3/models/user/request"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	repo request.Userlogin
}

func New(repo request.Userlogin) *AuthController {
	return &AuthController{
		repo: repo,
	}
}

func (ac *AuthController) Login() echo.HandlerFunc{
	return func(c echo.Context) error {
		
	}
}