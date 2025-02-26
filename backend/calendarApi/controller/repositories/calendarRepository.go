package repositories

import (
	"calendarApi/infrastructure/entity"
	"errors"
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

func (r *CalendarRepository) GetByDateId(date string, id string) (*entity.Calendar, error) {
	var entry *entity.Calendar
	if err := r.db.Where("date = ? AND id = ?", date, id).Find(&entry).Error; err != nil {
		return nil, err
	}
	return entry, nil
}

func (r *CalendarRepository) UpdateById(id string) (*entity.Calendar, error) {
	var calendar entity.Calendar
	if err := r.db.Where("id = ?", id).First(&calendar).Error; err != nil {
		return nil, err
	}
	return &calendar, nil
}

func (r *CalendarRepository) DeleteByID(id string) error {
	result := r.db.Delete(&entity.Calendar{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("calendar not found")
	}
	return nil
}
