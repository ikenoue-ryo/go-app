package main

import (
	"fmt"
	"go-project/go-app/app/controllers"
	"go-project/go-app/app/models"
)

func main() {
	fmt.Println(models.Db)

	controllers.StartMainServer()
}
