package controller

import (
	"calendarApi/controller/handler"
	"github.com/gin-gonic/gin"
)

func RouteController(calendarHandler *handler.CalendarHandler) {
	router := gin.Default()

	calendar := router.Group("/calendar")

	calendar.GET("/:date", calendarHandler.GetCalendarByDate)

	router.Run("localhost:8081")
}
