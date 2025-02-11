package repositories

import (
	"errors"
	"gorm.io/gorm"
	"math"
	"tracker/infrastructure/dto/model"
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

func (r *NutritionStatisticsRepository) AddStatistics(data model.ExpectedProductDTO) error {
	var stats entity.NutritionStatistics
	result := r.db.Where("date = ?", data.Date).First(&stats)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			// Create a new entry if none exists
			stats = entity.NutritionStatistics{
				Date:             data.Date,
				ConsumedCalories: 0,
				ConsumedProteins: 0,
				ConsumedFats:     0,
				ConsumedCarbs:    0,
			}
			r.db.Create(&stats)
		} else {
			return result.Error
		}
	}

	// Update the statistics
	stats.ConsumedCalories += data.Product.CaloriesPer100Grams * data.Weight / 100
	stats.ConsumedProteins += data.Product.ProteinsInGrams * float64(data.Weight) / 100
	stats.ConsumedFats += data.Product.FatsInGrams * float64(data.Weight) / 100
	stats.ConsumedCarbs += data.Product.CarbsInGrams * float64(data.Weight) / 100

	return r.db.Save(&stats).Error
}

func (r *NutritionStatisticsRepository) GetByID(id uint) (*entity.ConsumedProduct, error) {
	var product entity.ConsumedProduct
	result := r.db.First(&product, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}

func (r *NutritionStatisticsRepository) SubtractStatistics(consumedProduct *entity.ConsumedProduct, date string) error {
	var stats entity.NutritionStatistics
	result := r.db.Where("date = ?", date).First(&stats)

	if result.Error != nil {
		return result.Error
	}

	// Subtract the statistics with uint safety
	stats.ConsumedCalories = int(math.Max(float64(stats.ConsumedCalories-consumedProduct.Calories), 0))
	stats.ConsumedProteins = math.Max(stats.ConsumedProteins-consumedProduct.ProteinsInGrams, 0)
	stats.ConsumedFats = math.Max(stats.ConsumedFats-consumedProduct.FatsInGrams, 0)
	stats.ConsumedCarbs = math.Max(stats.ConsumedCarbs-consumedProduct.CarbsInGrams, 0)

	return r.db.Save(&stats).Error
}
