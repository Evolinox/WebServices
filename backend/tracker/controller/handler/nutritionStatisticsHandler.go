package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tracker/controller/helper"
	"tracker/controller/repositories"
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format in url. Use YYYY-MM-DD"})
		return
	}

	stats, err := h.repo.GetStatisticsByDate(dateStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, stats)
}
