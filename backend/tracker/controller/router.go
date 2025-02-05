package controller

import (
	"github.com/gin-gonic/gin"
	"tracker/controller/handler"
)

func RouteController(productHandler *handler.ProductHandler, settingsHandler *handler.SettingsHandler) {
	router := gin.Default()

	tracker := router.Group("/tracker")

	productRouter := tracker.Group("/products")
	productRouter.GET("/", productHandler.GetProducts)
	productRouter.GET("/:id", productHandler.GetProductByID)
	productRouter.GET("/name/:name", productHandler.GetProductByName)

	settingsRouter := tracker.Group("/settings")
	settingsRouter.GET("/", settingsHandler.GetSettings)
	settingsRouter.PATCH("/", settingsHandler.UpdateSettings)

	router.Run("localhost:8082")
}
