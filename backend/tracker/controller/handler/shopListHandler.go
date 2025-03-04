package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tracker/controller/helper"
	"tracker/controller/repositories"
	"tracker/infrastructure/dto/model"
)

type ShopListHandler struct {
	repo *repositories.ShopListAPIRepository
}

func NewShopListHandler(repo *repositories.ShopListAPIRepository) *ShopListHandler {
	return &ShopListHandler{repo: repo}
}

func (h *ShopListHandler) GetShoppingLists(c *gin.Context) {
	lists, err := h.repo.GetShoppingLists()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, lists)
}

func (h *ShopListHandler) GetShoppingListByID(c *gin.Context) {
	id := c.Param("id")
	list, err := h.repo.GetShoppingListByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if list == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Shopping list not found"})
		return
	}
	c.JSON(http.StatusOK, list)
}

func (h *ShopListHandler) CreateShoppingList(c *gin.Context) {
	var list model.ShoppingListDTO
	if err := c.ShouldBindJSON(&list); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !helper.IsValidDateFormat(list.Date) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format. Use YYYY-MM-DD"})
		return
	}

	err := h.repo.CreateShoppingList(list)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Shopping list created successfully"})
}

func (h *ShopListHandler) CreateShoppingListEntry(c *gin.Context) {
	id := c.Param("id")
	var entry model.ShoppingListEntryDTO
	if err := c.ShouldBindJSON(&entry); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.repo.CreateShoppingListEntry(id, entry)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Shopping list entry created successfully"})
}

func (h *ShopListHandler) DeleteShoppingList(c *gin.Context) {
	id := c.Param("id")

	err := h.repo.DeleteShoppingList(id)
	if err != nil {
		if err.Error() == "shopping list not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Shopping list not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Shopping list deleted successfully"})
}

func (h *ShopListHandler) DeleteShoppingListEntry(c *gin.Context) {
	id := c.Param("id")
	entryId := c.Param("entryID")

	err := h.repo.DeleteShoppingListEntry(id, entryId)
	if err != nil {
		if err.Error() == "shopping list entry not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Shopping list entry not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Shopping list entry deleted successfully"})
}
