package main

import (
	"fmt"
	"log"
	"part3/configs"
	"part3/delivery/controllers/auth"
	"part3/delivery/controllers/project"
	"part3/delivery/controllers/task"
	"part3/delivery/controllers/user"
	"part3/delivery/routes"
	_authDb "part3/lib/database/auth"
	_proDb "part3/lib/database/project"
	_taskDB "part3/lib/database/task"
	_userDb "part3/lib/database/user"
	"part3/utils"

	"github.com/labstack/echo/v4"
)

func main() {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	userRepo := _userDb.New(db)
	userController := user.New(userRepo)
	proRepo := _proDb.New(db)
	proController := project.NewRepo(proRepo)
	taskRepo := _taskDB.New(db)
	taskController := task.New(taskRepo,proRepo)
	authRepo := _authDb.New(db)
	authController := auth.New(authRepo)

	e := echo.New()

	routes.UserPath(e, userController, authController)
	routes.TaskPath(e, taskController)
	routes.ProjectPath(e, proController)
	routes.AdminPath(e, userController, authController)

	log.Fatal(e.Start(fmt.Sprintf(":%d", config.Port)))
}
