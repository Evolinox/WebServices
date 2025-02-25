package controller

import (
	"calendarApi/controller/handler"
	"github.com/gin-gonic/gin"
)

func RouteController(calendarHandler *handler.CalendarHandler) {
	router := gin.Default()

	calendar := router.Group("/calendar")

	calendar.GET("/:date", calendarHandler.GetCalendarByDate)
	calendar.POST("/", calendarHandler.CreateCalendar)
	calendar.PUT("/:id", calendarHandler.UpdateCalendar)
	calendar.DELETE("/:id", calendarHandler.DeleteCalendar)

	router.Run("localhost:8081")
}
