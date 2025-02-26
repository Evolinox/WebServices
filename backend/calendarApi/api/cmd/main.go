package main

import (
	"calendarApi/api"
	"calendarApi/config"
	"calendarApi/controller"
	"calendarApi/controller/handler"
)

func main() {
	config.LoadConfig("")
	dic := api.DIC{}
	calendarRepo := dic.GetCalendarRepository()
	calendarHandler := handler.NewCalendarHandler(calendarRepo)
	controller.RouteController(calendarHandler)
}
