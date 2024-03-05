package main

import (
	"github.com/gin-gonic/gin"

	"gin-project/controller"
)




func main() {

	personController := controller.NewPersonController()
	ginEngine := gin.Default()

	ginEngine.GET("/", rootHandler)
	ginEngine.GET("/person", personController.GetAll)
	ginEngine.POST("/person", personController.Create)

	err := ginEngine.Run("localhost:8082")

	if err != nil {
		panic(err)
	}
}


func rootHandler(ctx *gin.Context) {
	ctx.Writer.Write([]byte("Hello World"))
}