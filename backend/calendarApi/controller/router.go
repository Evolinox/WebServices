package controller

import (
	"calendarApi/controller/handler"
	"github.com/gin-gonic/gin"
)

func RouteController(calendarHandler *handler.CalendarHandler) {
	router := gin.Default()

	calendar := router.Group("/calendar")

	calendar.GET("/:date", calendarHandler.GetCalendarEntriesByDate)
	calendar.GET("/:date/:id", calendarHandler.GetCalendarEntryByDateId)
	calendar.POST("/", calendarHandler.CreateCalendarEntry)
	calendar.PATCH("/:id", calendarHandler.UpdateCalendarEntry)
	calendar.DELETE("/:id", calendarHandler.DeleteCalendarEntry)

	router.Run("localhost:8081")
}
