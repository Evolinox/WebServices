package handler

import "tracker/controller/repositories"

type CalendarHandler struct {
	repo *repositories.CalendarAPIRepository
}

func NewCalendarHandler(repo *repositories.CalendarAPIRepository) *CalendarHandler {
	return &CalendarHandler{repo: repo}
}
