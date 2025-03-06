package controller

import (
	"github.com/gin-gonic/gin"
	"productApi/controller/handler"
)

func RouteController(productHandler *handler.ProductHandler) {
	router := gin.Default()

	router.GET("/products", productHandler.GetProducts)
	router.GET("/products/:id", productHandler.GetProductByID)
	router.GET("/products/name/:name", productHandler.GetProductByName)
	router.POST("/products", productHandler.CreateProduct)
	router.DELETE("/products/:id", productHandler.DeleteProductByID)

	router.Run(":9402")
}
