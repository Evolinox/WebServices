package main

import (
	"tracker/api"
	"tracker/config"
	"tracker/controller"
	"tracker/controller/handler"
)

func main() {
	config.LoadConfig(".env")
	dic := api.DIC{}
	productAPIRepo := dic.GetProductAPIRepository()
	settingsRepo := dic.GetSettingsRepository()
	productHandler := handler.NewProductHandler(productAPIRepo)
	settingsHandler := handler.NewSettingsHandler(settingsRepo)

	controller.RouteController(productHandler, settingsHandler)
}
