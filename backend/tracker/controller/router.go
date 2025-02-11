package controller

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"tracker/controller/handler"
)

//	@title			Nutrition Tracker API
//	@version		1.0
//	@description	This is the Documentation of the Nutrition Tracker API, which contains 3 external APIs. The main Goal of the API is to get insights in your nutrition and helps you to organize your shopping and your meetings.

//	@contact.name	Tim Kerl
//	@contact.email	tim.kerl.23@lehre.mosbach.dhbw.de

//	@host		localhost:8082
//	@BasePath	/tracker/

//	@externalDocs.description	OpenAPI
//	@externalDocs.url			https://swagger.io/resources/open-api/

func RouteController(
	productHandler *handler.ProductHandler,
	consumeProductHandler *handler.ConsumeProductHandler,
	diaryHandler *handler.DiaryHandler,
	nutritionStatisticsHandler *handler.NutritionStatisticsHandler,
	settingsHandler *handler.SettingsHandler) {
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
	nutritionStaticsRouter.GET("/date/:date", nutritionStatisticsHandler.GetNutritionStatisticsByDate)

	settingsRouter := tracker.Group("/settings")
	settingsRouter.GET("/", settingsHandler.GetSettings)
	settingsRouter.PATCH("/", settingsHandler.UpdateSettings)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := router.Run("localhost:8082")
	if err != nil {
		return
	}
}
