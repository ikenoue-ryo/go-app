package main

import (
	"fmt"
	"go-project/go-app/app/controllers"
	"go-project/go-app/app/models"
)

func main() {
	fmt.Println(models.Db)

	u := &models.User{}
	u.Name = "test4"
	u.Email = "test4@example.com"
	u.PassWord = "testtest"
	fmt.Println(u)

	u.CreateUser()

	controllers.StartMainServer()
}
