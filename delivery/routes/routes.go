package routes

import (
	"part3/delivery/controllers/auth"
	"part3/delivery/controllers/project"
	"part3/delivery/controllers/task"
	"part3/delivery/controllers/user"
	"part3/delivery/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func UserPath(e *echo.Echo, uc *user.UserController, ac *auth.AuthController) {
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}",
	}))

	e.POST("/users", uc.Create())
	e.POST("/login", ac.Login())
	e.GET("/users/me", uc.GetById(), middlewares.JwtMiddleware())
	e.PUT("/users/me", uc.UpdateById(), middlewares.JwtMiddleware())
	e.DELETE("/users/me", uc.DeleteById(), middlewares.JwtMiddleware())
}

func TaskPath(e *echo.Echo, tc *task.TaskController) {
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}",
	}))

	// etask := e.Group("/todo",  middlewares.JwtMiddleware())
	e.POST("/todo/tasks", tc.Create(), middlewares.JwtMiddleware())
	e.GET("/todo/tasks", tc.GetAll(), middlewares.JwtMiddleware())
	e.PUT("/todo/tasks/:id", tc.Put(), middlewares.JwtMiddleware())
	e.PUT("/todo/tasks/:id", tc.UpdateByStatus(), middlewares.JwtMiddleware())
	e.DELETE("/todo/tasks/:id", tc.Delete(), middlewares.JwtMiddleware())
}

func ProjectPath(e *echo.Echo, pc *project.ProController) {
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}",
	}))

	e.POST("/projects", pc.Create(), middlewares.JwtMiddleware())
	e.GET("/projects", pc.GetAll(), middlewares.JwtMiddleware())
	e.PUT("/projects/:id", pc.Put(), middlewares.JwtMiddleware())
	e.DELETE("/projects/:id", pc.Delete(), middlewares.JwtMiddleware())
}

func AdminPath(e *echo.Echo, uc *user.UserController, ac *auth.AuthController) {
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}",
	}))

	e.GET("/admin/users", uc.GetAll(), middlewares.JwtMiddleware())
}
