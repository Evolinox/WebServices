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

	shoppingList.DELETE("/:id", shoppingListHandler.DeleteShoppingList)
	shoppingList.DELETE("/:id/products/:entryId", shoppingListHandler.DeleteShoppingListEntry) // Muss noch auf Funktion getestet werden :)

	err := router.Run(":8084")
	if err != nil {
		return
	}
}
