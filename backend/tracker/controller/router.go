package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
	router.Use(corsMiddleware())

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

	calendarRouter := tracker.Group("/calendar")
	calendarRouter.GET("/:date", calendarHandler.GetCalendarByDate)
	calendarRouter.POST("/", calendarHandler.AddCalendarEntry)
	calendarRouter.PATCH("/:id", calendarHandler.UpdateCalendarEntry)
	calendarRouter.DELETE("/:id", calendarHandler.DeleteCalendar)

	shopListRouter := tracker.Group("/shoppinglist")
	shopListRouter.GET("/", shopListHandler.GetShoppingLists)
	shopListRouter.GET("/:id", shopListHandler.GetShoppingListByID)
	shopListRouter.POST("/", shopListHandler.CreateShoppingList)
	shopListRouter.POST("/:id", shopListHandler.CreateShoppingListEntry)
	shopListRouter.DELETE("/:id", shopListHandler.DeleteShoppingList)
	shopListRouter.DELETE("/:id/products/:entryID", shopListHandler.DeleteShoppingListEntry)

	err := router.Run(":8082")
	if err != nil {
		return
	}
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, PATCH, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Writer.Header().Set("Access-Control-Max-Age", "600") // CORS-permission to cache data for 10 minutes

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}
