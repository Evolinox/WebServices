package handler

import (
	"calendarApi/controller/repositories"
	"calendarApi/infrastructure/entity"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
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

func (h *CalendarHandler) CreateCalendar(c *gin.Context) {
	var calendar entity.Calendar
	if err := c.ShouldBindJSON(&calendar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.repo.Create(&calendar)
	if err != nil {
		if strings.Contains(err.Error(), "1062") {
			c.JSON(http.StatusConflict, gin.H{"error": "Calendar ID already exists"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusCreated, calendar)
}

func (h *CalendarHandler) UpdateCalendar(c *gin.Context) {
	id := c.Param("id")
	calendar, err := h.repo.UpdateById(id)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, calendar)
}

func (h *CalendarHandler) DeleteCalendar(c *gin.Context) {
	id := c.Param("id")
	err := h.repo.DeleteByID(id)
	if err != nil {
		if err.Error() == "calendar not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Calendar not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "calendar deleted"})
}
