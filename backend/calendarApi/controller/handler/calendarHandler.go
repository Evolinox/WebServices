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

func NewCalendarHandler(repo *repositories.CalendarRepository) *CalendarHandler {
	return &CalendarHandler{repo: repo}
}

func (handler *CalendarHandler) GetCalendarEntriesByDate(c *gin.Context) {
	date := c.Param("date")
	product, err := handler.repo.GetByDate(date)
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

func (handler *CalendarHandler) GetCalendarEntryByDateId(context *gin.Context) {
	date := context.Param("date")
	id := context.Param("id")
	calendar, err := handler.repo.GetByDateId(date, id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if calendar == nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Date and ID not found"})
		return
	}

	context.JSON(http.StatusOK, calendar)
}

func (handler *CalendarHandler) CreateCalendarEntry(c *gin.Context) {
	var calendar entity.Calendar
	if err := c.ShouldBindJSON(&calendar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := handler.repo.Create(&calendar)
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

func (handler *CalendarHandler) UpdateCalendarEntry(c *gin.Context) {
	id := c.Param("id")
	var calendar entity.Calendar
	if err := c.ShouldBindJSON(&calendar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := handler.repo.UpdateById(id, &calendar)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, calendar)
}

func (handler *CalendarHandler) DeleteCalendarEntry(c *gin.Context) {
	id := c.Param("id")
	err := handler.repo.DeleteByID(id)
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
