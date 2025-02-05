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
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Create new settings if none exist
			return r.db.Create(updatedSettings).Error
		}
		return err
	}

	return r.db.Model(&settings).Updates(updatedSettings).Error
}
