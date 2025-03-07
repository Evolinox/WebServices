package main

import (
	"productApi/api"
	"productApi/config"
	"productApi/controller"
	"productApi/controller/handler"
)

func main() {
	config.LoadConfig("")
	dic := api.DIC{}
	productRepo := dic.GetProductRepository()
	productHandler := handler.NewProductHandler(productRepo)
	controller.RouteController(productHandler)
}
