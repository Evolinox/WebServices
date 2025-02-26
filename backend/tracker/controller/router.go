package controller

import (
	"github.com/gin-gonic/gin"
	"tracker/controller/handler"
)

func RouteController(
	productHandler *handler.ProductHandler,
	consumeProductHandler *handler.ConsumeProductHandler,
	diaryHandler *handler.DiaryHandler,
	nutritionStatisticsHandler *handler.NutritionStatisticsHandler,
	settingsHandler *handler.SettingsHandler,
	shopListHandler *handler.ShopListHandler,
	calendarHandler *handler.CalendarHandler) {
	router := gin.Default()

	tracker := router.Group("/tracker")

	productRouter := tracker.Group("/products")
	productRouter.GET("/", productHandler.GetProducts)
	productRouter.GET("/:id", productHandler.GetProductByID)
	productRouter.GET("/name/:name", productHandler.GetProductByName)

	trackProductRouter := tracker.Group("/consume")
	trackProductRouter.POST("/", consumeProductHandler.ConsumeProduct)

	diaryRouter := tracker.Group("/diary/date")
	diaryRouter.GET("/:date", diaryHandler.GetDiaryByDate)
	diaryRouter.GET("/:date/:id", diaryHandler.GetConsumedProductByID)
	diaryRouter.DELETE("/:date/:id", consumeProductHandler.DeleteConsumedProduct)

	nutritionStaticsRouter := tracker.Group("/nutrition")
	nutritionStaticsRouter.GET("/:date", nutritionStatisticsHandler.GetNutritionStatisticsByDate)

	settingsRouter := tracker.Group("/settings")
	settingsRouter.GET("/", settingsHandler.GetSettings)
	settingsRouter.PATCH("/", settingsHandler.UpdateSettings)

	err := router.Run(":8082")
	if err != nil {
		return
	}
}
