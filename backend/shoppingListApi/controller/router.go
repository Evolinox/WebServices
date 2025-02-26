package controller

import (
	"github.com/gin-gonic/gin"
	"shoppingListApi/controller/handler"
)

func RouteController(shoppingListHandler *handler.ShoppingListHandler) {
	router := gin.Default()

	shoppinglist := router.Group("/shoppinglist")

	router.Run("localhost:8081")
}
