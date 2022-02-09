package main

import (
	"fmt"
	"log"
	"part3/configs"
	"part3/delivery/controllers/auth"
	"part3/delivery/controllers/user"
	"part3/delivery/routes"
	_authDb "part3/lib/database/auth"
	_userDb "part3/lib/database/user"
	"part3/utils"

	"github.com/labstack/echo/v4"
)

func main() {
	config := configs.GetConfig()

	db := utils.InitDB(config)

	userRepo := _userDb.New(db)
	userController := user.New(userRepo)
	authRepo := _authDb.New(db)
	authController := auth.New(authRepo)

	e := echo.New()

	routes.UserPath(e, userController, authController)

	log.Fatal(e.Start(fmt.Sprintf(":%d", config.Port)))
}
