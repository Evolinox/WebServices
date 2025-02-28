package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shoppingListApi/controller/repositories"
	"shoppingListApi/infrastructure/entity"
	"strings"
)

type ShoppingListHandler struct {
	repo *repositories.ShoppingListRepository
}

func NewShoppingListHandler(repo *repositories.ShoppingListRepository) *ShoppingListHandler {
	return &ShoppingListHandler{repo: repo}
}

func (handler *ShoppingListHandler) CreateShoppingList(context *gin.Context) {
	var list entity.ShoppingList
	if err := context.ShouldBindJSON(&list); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := handler.repo.Create(&list)
	if err != nil {
		if strings.Contains(err.Error(), "1062") {
			context.JSON(http.StatusConflict, gin.H{"error": "ShoppingList already exists"})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	context.JSON(http.StatusCreated, list)
}

func (handler *ShoppingListHandler) CreateShoppingListEntry(context *gin.Context) {
	id := context.Param("id")
	var product entity.Product
	if err := context.ShouldBindJSON(&product); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := handler.repo.CreateProduct(id, &product)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, product)

}

func (handler *ShoppingListHandler) GetShoppingLists(context *gin.Context) {
	lists, err := handler.repo.GetAll()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if lists == nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "No shopping lists for that date found"})
		return
	}
	context.JSON(http.StatusOK, lists)
}

func (handler *ShoppingListHandler) GetShoppingListByDate(context *gin.Context) {
	date := context.Param("date")
	list, err := handler.repo.GetByDate(date)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if list == nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "No shopping lists for that date found"})
		return
	}
	context.JSON(http.StatusOK, list)
}

func (handler *ShoppingListHandler) GetShoppingListById(context *gin.Context) {
	id := context.Param("id")
	list, err := handler.repo.GetById(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if list == nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "No shopping list with that id found"})
		return
	}
	context.JSON(http.StatusOK, list)
}

func (handler *ShoppingListHandler) GetShoppingListByDateById(context *gin.Context) {
	date := context.Param("date")
	id := context.Param("id")
	list, err := handler.repo.GetByDateById(date, id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if list == nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "No shopping list for this date and id found"})
		return
	}
	context.JSON(http.StatusOK, list)
}

func (handler *ShoppingListHandler) UpdateShoppingList(context *gin.Context) {
	id := context.Param("id")
	var list entity.ShoppingList
	if err := context.ShouldBindJSON(&list); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := handler.repo.UpdateById(id, &list)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, list)
}

func (handler *ShoppingListHandler) DeleteShoppingList(context *gin.Context) {
	id := context.Param("id")
	err := handler.repo.DeleteById(id)
	if err != nil {
		if err.Error() == "shoppinglist not found" {
			context.JSON(http.StatusNotFound, gin.H{"error": "Shoppinglist not found"})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "calendar deleted"})
}
