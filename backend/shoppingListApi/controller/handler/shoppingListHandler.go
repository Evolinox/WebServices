package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shoppingListApi/controller/repositories"
	"shoppingListApi/infrastructure/entity"
	"strconv"
)

type ShoppingListHandler struct {
	repo *repositories.ShoppingListRepository
}

func NewShoppingListHandler(repo *repositories.ShoppingListRepository) *ShoppingListHandler {
	return &ShoppingListHandler{repo: repo}
}

func (h *ShoppingListHandler) GetShoppingLists(c *gin.Context) {
	lists, err := h.repo.GetShoppingLists()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch shopping lists"})
		return
	}
	c.JSON(http.StatusOK, lists)
}

func (h *ShoppingListHandler) GetShoppingListById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid shopping list ID"})
		return
	}
	list, err := h.repo.GetShoppingListById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Shopping list not found"})
		return
	}
	c.JSON(http.StatusOK, list)
}

func (h *ShoppingListHandler) CreateShoppingList(c *gin.Context) {
	var list entity.ShoppingList
	if err := c.ShouldBindJSON(&list); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := h.repo.CreateShoppingList(&list); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create shopping list"})
		return
	}
	c.JSON(http.StatusCreated, list)
}

func (h *ShoppingListHandler) CreateShoppingListEntry(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid shopping list ID"})
		return
	}
	var product entity.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	product.ShoppingListID = id
	if err := h.repo.CreateShoppingListEntry(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add product"})
		return
	}
	c.JSON(http.StatusCreated, product)
}

func (h *ShoppingListHandler) DeleteShoppingList(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid shopping list ID"})
		return
	}
	if err := h.repo.DeleteShoppingList(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete shopping list"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Shopping list deleted"})
}

func (h *ShoppingListHandler) DeleteShoppingListEntry(c *gin.Context) {
	entryId, err := strconv.Atoi(c.Param("entryId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}
	if err := h.repo.DeleteShoppingListEntry(entryId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}
