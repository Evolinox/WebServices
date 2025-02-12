package handler

import (
	"calendarApi/controller/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CalendarHandler struct {
	repo *repositories.CalendarRepository
}

func (h *CalendarHandler) GetCalendarByDate(c *gin.Context) {
	date := c.Param("date")
	product, err := h.repo.GetByDate(date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if product == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Date not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}
