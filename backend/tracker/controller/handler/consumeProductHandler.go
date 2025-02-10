package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"tracker/controller/repositories"
	"tracker/infrastructure/dto/mapper"
	"tracker/infrastructure/dto/model"
)

type ConsumeProductHandler struct {
	repo *repositories.ConsumeProductRepository
}

func NewConsumeProductHandler(repo *repositories.ConsumeProductRepository) *ConsumeProductHandler {
	return &ConsumeProductHandler{repo: repo}
}

func (h *ConsumeProductHandler) ConsumeProduct(c *gin.Context) {
	var receivedProductData model.ReceiveProductDTO
	product := mapper.CalculateNutritionValues(receivedProductData)
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.repo.Create(&product)
	if err != nil {
		if strings.Contains(err.Error(), "1062") {
			c.JSON(http.StatusConflict, gin.H{"error": "Product ID already exists"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusCreated, product)
}

func (h *ConsumeProductHandler) DeleteConsumedProduct(c *gin.Context) {
	id := c.Param("id")
	err := h.repo.DeleteByID(id)
	if err != nil {
		if err.Error() == "product not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
