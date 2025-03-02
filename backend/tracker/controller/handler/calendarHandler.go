package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
	entries, err := h.repo.GetCalendarByDate(date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if entries == nil || len(entries) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No calendar entries found for the given date"})
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
