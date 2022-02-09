package main

import (
	"fmt"
	"part3/configs"
	"part3/route"
	"part3/utils"
)

func main() {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	fmt.Print(db)

	e := route.New()
	e.Start(":8000")

}
