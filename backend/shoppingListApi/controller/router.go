package controller

import (
	"github.com/gin-gonic/gin"
	"shoppingListApi/controller/handler"
)

func RouteController(shoppingListHandler *handler.ShoppingListHandler) {
	router := gin.Default()

	shoppingList := router.Group("/shoppinglist")

	shoppingList.GET("/", shoppingListHandler.GetShoppingLists)
	shoppingList.GET("/:id", shoppingListHandler.GetShoppingListById)

	shoppingList.POST("/", shoppingListHandler.CreateShoppingList)
	shoppingList.POST("/:id")

	shoppingList.PATCH("/:id", shoppingListHandler.UpdateShoppingList)
	shoppingList.PATCH("/:id/:entry")

	shoppingList.DELETE("/:id", shoppingListHandler.DeleteShoppingList)
	shoppingList.DELETE("/:id/:entry")

	router.Run("localhost:8084")
}
