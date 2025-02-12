package repositories

import (
	"errors"
	"gorm.io/gorm"
	"tracker/infrastructure/entity"
)

type DiaryRepository struct {
	db *gorm.DB
}

func NewDiaryRepository(db *gorm.DB) *DiaryRepository {
	return &DiaryRepository{db: db}
}

func (r *DiaryRepository) GetDiaryByDate(date string) (*entity.DailyProductsConsumed, error) {
	var dailyRecord entity.DailyProductsConsumed

	result := r.db.Preload("Products").Where("DATE(date) = ?", date).First(&dailyRecord)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			newRecord := entity.DailyProductsConsumed{
				Date:     date,
				Products: []entity.ConsumedProduct{}, // No products yet
			}
			if err := r.db.Create(&newRecord).Error; err != nil {
				return nil, err
			}
			return &newRecord, nil
		}
		return nil, result.Error
	}

	return &dailyRecord, nil
}

func (r *DiaryRepository) GetConsumedProductByID(id uint) (*entity.ConsumedProduct, error) {
	var dailyRecords []entity.DailyProductsConsumed

	result := r.db.Preload("Products").Find(&dailyRecords)
	if result.Error != nil {
		return nil, errors.New("failed to fetch daily records")
	}

	// Search for the product inside the Products array
	for _, dailyRecord := range dailyRecords {
		for _, product := range dailyRecord.Products {
			if product.ID == id {
				return &product, nil
			}
		}
	}

	// If product not found, return an error
	return nil, errors.New("product not found")
}
