package main

import (
	"rest_api/controller"
	"rest_api/db"
	"rest_api/repository"
	"rest_api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	ItemRepository := repository.NewItemRepository(dbConnection)
	ItemUsecase := usecase.NewItemUsecase(ItemRepository)
	ItemController := controller.NewItemController(ItemUsecase)

	server.GET("/items", ItemController.GetItems)
	server.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Successfully Tested",
		})
	})

	server.Run(":8000")
}
