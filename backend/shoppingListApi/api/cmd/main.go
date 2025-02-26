package main

import (
	"shoppingListApi/api"
	"shoppingListApi/config"
	"shoppingListApi/controller"
	"shoppingListApi/controller/handler"
)

func main() {
	config.LoadConfig("")
	dic := api.DIC{}
	shoppingListRepo := dic.GetShoppingListRepository()
	shoppingListHandler := handler.NewShoppingListHandler(shoppingListRepo)
	controller.RouteController(shoppingListHandler)
}
