package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tracker/controller/helper"
	"tracker/controller/repositories"
	"tracker/infrastructure/dto/model"
)

type CalendarHandler struct {
	repo *repositories.CalendarAPIRepository
}

func NewCalendarHandler(repo *repositories.CalendarAPIRepository) *CalendarHandler {
	return &CalendarHandler{repo: repo}
}

func (h *CalendarHandler) GetCalendarByDate(c *gin.Context) {
	date := c.Param("date")
	if !helper.IsValidDateFormat(date) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format in url. Use YYYY-MM-DD"})
		return
	}
	entries, err := h.repo.GetCalendarByDate(date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, entries)
}

func (h *CalendarHandler) AddCalendarEntry(c *gin.Context) {
	var entry model.CalendarDTO
	if err := c.ShouldBindJSON(&entry); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !helper.IsValidDateFormat(entry.Date) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format. Use YYYY-MM-DD"})
		return
	}
	if !helper.IsValidTimeFormat(entry.BeginTime) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid time format. Use HH-TT"})
		return
	}
	if !helper.IsValidTimeFormat(entry.EndTime) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid time format. Use HH-TT"})
		return
	}

	err := h.repo.AddCalendarEntry(entry)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Calendar entry added successfully"})
}

func (h *CalendarHandler) UpdateCalendarEntry(c *gin.Context) {
	id := c.Param("id")
	var entry model.CalendarDTO
	if err := c.ShouldBindJSON(&entry); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !helper.IsValidDateFormat(entry.Date) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format. Use YYYY-MM-DD"})
		return
	}
	if !helper.IsValidTimeFormat(entry.BeginTime) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid time format. Use HH-TT"})
		return
	}
	if !helper.IsValidTimeFormat(entry.EndTime) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid time format. Use HH-TT"})
		return
	}

	err := h.repo.UpdateCalendarEntry(id, entry)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Calendar entry updated successfully"})
}

func (h *CalendarHandler) DeleteCalendar(c *gin.Context) {
	id := c.Param("id")

	err := h.repo.DeleteCalendar(id)
	if err != nil {
		if err.Error() == "calendar entry not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Calendar entry not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Calendar entry deleted successfully"})
}
