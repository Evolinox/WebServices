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
	shoppingList.POST("/:id", shoppingListHandler.CreateShoppingListEntry) // Muss noch auf Funktion getestet werden :)

	shoppingList.PATCH("/:id", shoppingListHandler.UpdateShoppingList)
	shoppingList.PATCH("/:id/products/:entryId", shoppingListHandler.UpdateShoppingListEntry) // Muss noch auf Funktion getestet werden :)

	shoppingList.DELETE("/:id", shoppingListHandler.DeleteShoppingList)
	shoppingList.DELETE("/:id/products/:entryId", shoppingListHandler.DeleteShoppingListEntry) // Muss noch auf Funktion getestet werden :)

	router.Run("localhost:8084")
}
