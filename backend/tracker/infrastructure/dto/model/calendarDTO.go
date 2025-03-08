package model

type CalendarDTO struct {
	ID          int    `json:"ID"`
	Date        string `json:"Date"`
	BeginTime   string `json:"BeginTime"`
	EndTime     string `json:"EndTime"`
	Name        string `json:"Name"`
	Description string `json:"Description"`
}
