package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"tracker/controller/helper"
	"tracker/controller/repositories"
	"tracker/infrastructure/dto/model"
)

type NutritionStatisticsHandler struct {
	repo *repositories.NutritionStatisticsRepository
}

func NewNutritionStatisticsHandler(repo *repositories.NutritionStatisticsRepository) *NutritionStatisticsHandler {
	return &NutritionStatisticsHandler{repo: repo}
}

func (h *NutritionStatisticsHandler) GetNutritionStatisticsByDate(c *gin.Context) {
	dateStr := c.Param("date")

	if !helper.IsValidDateFormat(dateStr) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format. Use YYYY-MM-DD"})
		return
	}

	stats, err := h.repo.GetStatisticsByDate(dateStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, stats)
}

func (h *NutritionStatisticsHandler) AddNutritionStatistics(c *gin.Context) {
	var receivedProductData model.ExpectedProductDTO

	if err := c.ShouldBindJSON(&receivedProductData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.repo.AddStatistics(receivedProductData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update nutrition statistics"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Nutrition statistics updated successfully"})
}

func (h *NutritionStatisticsHandler) SubtractNutritionStatistics(c *gin.Context) {
	id := c.Param("id")
	date := c.Param("date")

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

	// Subtract the nutrition statistics
	err = h.repo.SubtractStatistics(consumedProduct, date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to subtract nutrition statistics"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Nutrition statistics subtracted successfully"})
}
