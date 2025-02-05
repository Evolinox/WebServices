package main

import (
	"tracker/api"
	"tracker/controller"
	"tracker/controller/handler"
)

func main() {
	dic := api.DIC{}
	productAPIRepo := dic.GetProductAPIRepository()
	productHandler := handler.NewProductHandler(productAPIRepo)

	controller.RouteController(productHandler)
}
