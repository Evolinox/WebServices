package repositories

type CalendarAPIRepository struct {
	baseURL string
}

func NewCalendarAPIRepository(baseURL string) *CalendarAPIRepository {
	return &CalendarAPIRepository{baseURL: baseURL}
}
