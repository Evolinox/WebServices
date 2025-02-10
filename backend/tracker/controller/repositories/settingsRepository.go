package repositories

import (
	"errors"
	"gorm.io/gorm"
	"tracker/infrastructure/entity"
)

type SettingsRepository struct {
	db *gorm.DB
}

func NewSettingsRepository(db *gorm.DB) *SettingsRepository {
	return &SettingsRepository{db: db}
}

func (r *SettingsRepository) GetSettings() (*entity.Settings, error) {
	var settings entity.Settings
	if err := r.db.First(&settings).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // No settings found
		}
		return nil, err
	}
	return &settings, nil
}

func (r *SettingsRepository) UpdateSettings(updatedSettings *entity.Settings) error {
	var settings entity.Settings
	if err := r.db.First(&settings).Error; err != nil {
		return err
	}
	//Query 1=1 is used because there is only one record in the table
	return r.db.Model(&entity.Settings{}).Where("1 = 1").Updates(updatedSettings).Error
}
