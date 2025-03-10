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
	shoppingList.POST("/:id", shoppingListHandler.CreateShoppingListEntry)

	shoppingList.DELETE("/:id", shoppingListHandler.DeleteShoppingList)
	shoppingList.DELETE("/:id/products/:entryId", shoppingListHandler.DeleteShoppingListEntry)

	err := router.Run(":9404")
	if err != nil {
		return
	}
}
