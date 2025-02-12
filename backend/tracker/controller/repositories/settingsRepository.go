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

	// Perform the update (using 1=1 to ensure it's the only record)
	return r.db.Model(&entity.Settings{}).Where("1 = 1").Updates(updatedSettings).Error
}
