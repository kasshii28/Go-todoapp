package main

import (
	"fmt"
	"Go-todoapp/app/models"
	"Go-todoapp/app/models/controllers"
)

func main() {
	fmt.Println(models.Db)

	controllers.StartMainServer()
	
}
