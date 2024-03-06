package main

import (
	"gin-person/controller"
	"gin-person/lib"
	"gin-person/repository"

	"github.com/gin-gonic/gin"
)

func main() {

	db, err := lib.InitializeDatabase()
	if err != nil {
		panic(err)
	}

	personRepository := repository.NewPersonRepository(db)
	personController := controller.NewPersonController(personRepository)

	ginEngine := gin.Default()

	ginEngine.GET("/person", personController.GetAll)
	ginEngine.POST("/person", personController.Create)

	err = ginEngine.Run("localhost:8082")
	if err != nil {
		panic(err)
	}
}