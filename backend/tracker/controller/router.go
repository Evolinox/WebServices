package controller

import (
	"github.com/gin-gonic/gin"
	"tracker/controller/handler"
)

func RouteController(productHandler *handler.ProductHandler) {
	router := gin.Default()

	tracker := router.Group("/tracker")

	productRouter := tracker.Group("/products")
	productRouter.GET("/", productHandler.GetProducts)
	productRouter.GET("/:id", productHandler.GetProductByID)
	productRouter.GET("/name/:name", productHandler.GetProductByName)

	router.Run("localhost:8082")
}
