package routes

import (
	"part3/delivery/controllers/auth"
	"part3/delivery/controllers/user"
	"part3/delivery/middlewares"

	"github.com/labstack/echo/v4"
)

func UserPath(e *echo.Echo, uc *user.UserController, ac *auth.AuthController) {
	e.POST("/users", uc.Create())
	e.POST("/login", ac.Login())
	e.GET("/users/me", uc.GetById(), middlewares.JwtMiddleware())
	e.PUT("/users/me", uc.UpdateById(),  middlewares.JwtMiddleware())
	e.DELETE("/users/me", uc.DeleteById(),  middlewares.JwtMiddleware())
}

