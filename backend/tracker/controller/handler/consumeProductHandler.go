package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
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

	// Bind JSON request body
	if err := c.ShouldBindJSON(&receivedProductData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get or create the corresponding DailyProductsConsumed entry
	dailyRecord, err := h.repo.GetOrCreateDailyRecord(receivedProductData.Date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find or create daily record"})
		return
	}

	// Use mapper to calculate nutrition values and assign daily record ID
	consumedProduct := mapper.CalculateNutritionValues(receivedProductData, dailyRecord.ID)

	// Save the consumed product
	err = h.repo.Create(consumedProduct)
	if err != nil {
		if strings.Contains(err.Error(), "1062") {
			c.JSON(http.StatusConflict, gin.H{"error": "Product ID already exists"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusCreated, consumedProduct)
}

func (h *ConsumeProductHandler) DeleteConsumedProduct(c *gin.Context) {
	id := c.Param("id")

	// Convert ID to uint
	productID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID format"})
		return
	}

	// Retrieve the consumed product before deleting
	consumedProduct, err := h.repo.GetByID(uint(productID))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	// Delete the consumed product
	err = h.repo.DeleteByID(uint(productID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete consumed product"})
		return
	}

	// Check if the DailyProductsConsumed entry is now empty
	err = h.repo.CleanupDailyRecord(consumedProduct.DailyProductsConsumedID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to clean up daily record"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
