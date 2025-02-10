package repositories

import (
	"errors"
	"gorm.io/gorm"
	"tracker/infrastructure/entity"
)

type NutritionStatisticsRepository struct {
	db *gorm.DB
}

func NewNutritionStatisticsRepository(db *gorm.DB) *NutritionStatisticsRepository {
	return &NutritionStatisticsRepository{db: db}
}

func (r *NutritionStatisticsRepository) GetStatisticsByDate(date string) (*entity.NutritionStatistics, error) {
	var stats entity.NutritionStatistics
	result := r.db.Where("date = ?", date).First(&stats)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			newStats := entity.NutritionStatistics{
				ID:               0,
				Date:             date,
				ConsumedCalories: 0,
				ConsumedProteins: 0,
				ConsumedFats:     0,
				ConsumedCarbs:    0,
			}
			if err := r.db.Create(&newStats).Error; err != nil {
				return nil, err
			}
			return &newStats, nil
		}
		return nil, result.Error
	}

	return &stats, nil
}
