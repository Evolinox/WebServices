package controller

import (
	"calendarApi/controller/handler"
	"github.com/gin-gonic/gin"
)

func RouteController(calendarHandler *handler.CalendarHandler) {
	router := gin.Default()

	router.GET("/date/:date", calendarHandler.GetCalendarByDate)

	router.Run("localhost:8081")
}
