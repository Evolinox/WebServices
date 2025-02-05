package controller

import (
	"github.com/gin-gonic/gin"
	"tracker/controller/handler"
)

func RouteController(productHandler *handler.ProductHandler) {
	router := gin.Default()

	router.GET("/tracker/products", productHandler.GetProducts)
	router.GET("/tracker/products/:id", productHandler.GetProductByID)
	router.GET("/tracker/products/name/:name", productHandler.GetProductByName)

	router.Run("localhost:8082")
}
