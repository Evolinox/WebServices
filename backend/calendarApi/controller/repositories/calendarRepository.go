package repositories

import (
	"calendarApi/infrastructure/entity"
	"gorm.io/gorm"
)

type CalendarRepository struct {
	db *gorm.DB
}

func NewCalendarRepository(db *gorm.DB) *CalendarRepository {
	return &CalendarRepository{db: db}
}

func (r *CalendarRepository) Create(entry *entity.Calendar) error {
	if err := r.db.Create(entry).Error; err != nil {
		return err
	}
	return nil
}

func (r *CalendarRepository) GetByDate(date string) ([]*entity.Calendar, error) {
	var entries []*entity.Calendar
	if err := r.db.Where("date = ?", date).Find(&entries).Error; err != nil {
		return nil, err
	}
	return entries, nil
}
