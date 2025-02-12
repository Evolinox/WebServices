package repositories

import (
	"errors"
	"gorm.io/gorm"
	"math"
	"tracker/infrastructure/entity"
)

// ToDO Update funzt net
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

func (r *NutritionStatisticsRepository) AddStatistics(consumedProduct *entity.ConsumedProduct, date string) error {
	var stats entity.NutritionStatistics

	result := r.db.Where("date = ?", date).First(&stats)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			stats = entity.NutritionStatistics{
				Date:             date,
				ConsumedCalories: 0,
				ConsumedProteins: 0,
				ConsumedFats:     0,
				ConsumedCarbs:    0,
			}
			if err := r.db.Create(&stats).Error; err != nil {
				return err
			}
			r.db.Where("date = ?", date).First(&stats)
		} else {
			return result.Error
		}
	}

	stats.ConsumedCalories += consumedProduct.Calories
	stats.ConsumedProteins += consumedProduct.ProteinsInGrams
	stats.ConsumedFats += consumedProduct.FatsInGrams
	stats.ConsumedCarbs += consumedProduct.CarbsInGrams

	return r.db.Save(&stats).Error
}

func (r *NutritionStatisticsRepository) SubtractStatistics(consumedProduct *entity.ConsumedProduct, date string) error {
	var stats entity.NutritionStatistics
	result := r.db.Where("date = ?", date).First(&stats)

	if result.Error != nil {
		return result.Error
	}

	stats.ConsumedCalories = int(math.Max(float64(stats.ConsumedCalories-consumedProduct.Calories), 0))
	stats.ConsumedProteins = math.Max(stats.ConsumedProteins-consumedProduct.ProteinsInGrams, 0)
	stats.ConsumedFats = math.Max(stats.ConsumedFats-consumedProduct.FatsInGrams, 0)
	stats.ConsumedCarbs = math.Max(stats.ConsumedCarbs-consumedProduct.CarbsInGrams, 0)

	return r.db.Save(&stats).Error
}
