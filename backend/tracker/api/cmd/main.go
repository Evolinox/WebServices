package main

import (
	"tracker/api"
	"tracker/controller"
	"tracker/controller/handler"
)

func main() {
	dic := api.DIC{}
	productAPIRepo := dic.GetProductAPIRepository()
	settingsRepo := dic.GetSettingsRepository()
	productHandler := handler.NewProductHandler(productAPIRepo)
	settingsHandler := handler.NewSettingsHandler(settingsRepo)

	controller.RouteController(productHandler, settingsHandler)
}
