package main

import (
	"Go-todoapp/app/models"
	"Go-todoapp/app/models/controllers"
	"fmt"
)

func main() {
	fmt.Println(models.Db)

	controllers.StartMainServer()
}