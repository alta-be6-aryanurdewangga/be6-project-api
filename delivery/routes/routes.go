package routes

import (
	"part3/delivery/controllers/auth"
	"part3/delivery/controllers/user"

	"github.com/labstack/echo/v4"
)

func RegisterPath(e *echo.Echo, uc *user.UserController, ac *auth.AuthController) {
	e.POST("/users", uc.Create())
	
}