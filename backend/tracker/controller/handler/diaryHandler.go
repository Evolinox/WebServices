package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"tracker/controller/helper"
	"tracker/controller/repositories"
)

type DiaryHandler struct {
	repo *repositories.DiaryRepository
}

func NewDiaryHandler(repo *repositories.DiaryRepository) *DiaryHandler {
	return &DiaryHandler{repo: repo}
}

func (h *DiaryHandler) GetDiaryByDate(c *gin.Context) {
	dateStr := c.Param("date") // e.g., "2024-02-09"

	if !helper.IsValidDateFormat(dateStr) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format. Use YYYY-MM-DD"})
		return
	}

	// Call repository function to fetch data
	dailyRecord, err := h.repo.GetDiaryByDate(dateStr)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "No records found for this date"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve diary entry"})
		return
	}

	// Return JSON response
	c.JSON(http.StatusOK, gin.H{
		"date":     dailyRecord.Date, // Keep as string (no need for formatting)
		"products": dailyRecord.Products,
	})
}

func (h *DiaryHandler) GetConsumedProductByID(c *gin.Context) {
	// Extract product ID from URL
	idStr := c.Param("id")

	// Convert ID string to integer
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format. Must be an integer"})
		return
	}

	consumedProduct, err := h.repo.GetConsumedProductByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Return JSON response
	c.JSON(http.StatusOK, consumedProduct)

}
