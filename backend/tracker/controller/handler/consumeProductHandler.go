package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"tracker/controller/repositories"
	"tracker/infrastructure/dto/mapper"
	"tracker/infrastructure/dto/model"
)

type ConsumeProductHandler struct {
	consumeRepo   *repositories.ConsumeProductRepository
	nutritionRepo *repositories.NutritionStatisticsRepository
}

func NewConsumeProductHandler(consumeRepo *repositories.ConsumeProductRepository, nutritionRepo *repositories.NutritionStatisticsRepository) *ConsumeProductHandler {
	return &ConsumeProductHandler{
		consumeRepo:   consumeRepo,
		nutritionRepo: nutritionRepo,
	}
}

func (h *ConsumeProductHandler) ConsumeProduct(c *gin.Context) {
	var receivedProductData model.ExpectedProductDTO

	// Bind JSON request data
	if err := c.ShouldBindJSON(&receivedProductData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get or create the daily record
	dailyRecordID, err := h.consumeRepo.GetOrCreateDailyRecord(receivedProductData.Date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create/retrieve daily record"})
		return
	}

	// Calculate the nutrition values
	consumedProduct := mapper.CalculateNutritionValues(receivedProductData, dailyRecordID)

	// Save the consumed product
	if err := h.consumeRepo.ConsumeProduct(consumedProduct); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to consume product"})
		return
	}

	// Update nutrition statistics
	if err := h.nutritionRepo.AddStatistics(consumedProduct, receivedProductData.Date); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update nutrition statistics"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":         "Product consumed and nutrition statistics updated successfully",
		"consumedProduct": consumedProduct,
	})
}

func (h *ConsumeProductHandler) DeleteConsumedProduct(c *gin.Context) {
	id := c.Param("id")
	date := c.Param("date")

	// Convert ID to uint
	productID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID format"})
		return
	}

	consumedProduct, err := h.consumeRepo.GetByID(uint(productID))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	if err := h.consumeRepo.DeleteConsumedProduct(uint(productID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}

	if err := h.nutritionRepo.SubtractStatistics(consumedProduct, date); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update nutrition statistics"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted and nutrition statistics updated successfully"})
}
