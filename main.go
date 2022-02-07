package main

import (
	"fmt"
	"part3/configs"
	"part3/utils"
)

func main() {
	config := configs.GetConfig()

	db := utils.InitDB(config)

	fmt.Print(db)

}
