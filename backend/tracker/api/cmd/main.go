package main

import (
	"tracker/api"
	"tracker/config"
	"tracker/controller"
	"tracker/controller/handler"
)

func main() {
	config.LoadConfig("")
	dic := api.DIC{}

	consumeProductRepo := dic.GetConsumeProductRepository()
	nutritionStatsRepo := dic.GetNutritionStatisticsRepository()
	diaryRepo := dic.GetDiaryRepository()
	settingsRepo := dic.GetSettingsRepository()
	productAPIRepo := dic.GetProductAPIRepository()
	shopListAPIRepo := dic.GetShopListAPIRepository()
	calendarAPIRepo := dic.GetCalendarAPIRepository()

	consumeProductHandler := handler.NewConsumeProductHandler(consumeProductRepo, nutritionStatsRepo)
	nutritionStatisticsHandler := handler.NewNutritionStatisticsHandler(nutritionStatsRepo)
	diaryHandler := handler.NewDiaryHandler(diaryRepo)
	settingsHandler := handler.NewSettingsHandler(settingsRepo)
	productHandler := handler.NewProductHandler(productAPIRepo)
	shopListHandler := handler.NewShopListHandler(shopListAPIRepo)
	calendarHandler := handler.NewCalendarHandler(calendarAPIRepo)

	controller.RouteController(
		productHandler,
		consumeProductHandler,
		diaryHandler,
		nutritionStatisticsHandler,
		settingsHandler,
		shopListHandler,
		calendarHandler,
	)
}
