package controller

import (
	"github.com/gin-gonic/gin"
	"shoppingListApi/controller/handler"
)

func RouteController(shoppingListHandler *handler.ShoppingListHandler) {
	router := gin.Default()

	shoppingList := router.Group("/shoppinglist")

	shoppingList.GET("/:date", shoppingListHandler.GetShoppingListByDate)
	shoppingList.GET("/:date/:id", shoppingListHandler.GetShoppingListByDateById)
	shoppingList.GET("/:id", shoppingListHandler.GetShoppingListById)
	shoppingList.POST("/", shoppingListHandler.CreateShoppingList)
	shoppingList.PATCH("/:id", shoppingListHandler.UpdateShoppingList)
	shoppingList.DELETE("/:id", shoppingListHandler.DeleteShoppingList)

	router.Run("localhost:8084")
}
