package repositories

import (
	"errors"
	"gorm.io/gorm"
	"tracker/infrastructure/entity"
)

type ConsumeProductRepository struct {
	db *gorm.DB
}

func NewConsumeProductRepository(db *gorm.DB) *ConsumeProductRepository {
	return &ConsumeProductRepository{db: db}
}

func (r *ConsumeProductRepository) GetOrCreateDailyRecord(date string) (*entity.DailyProductsConsumed, error) {
	var dailyRecord entity.DailyProductsConsumed

	result := r.db.Where("date = ?", date).First(&dailyRecord)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			dailyRecord = entity.DailyProductsConsumed{Date: date}
			if err := r.db.Create(&dailyRecord).Error; err != nil {
				return nil, err
			}
		} else {
			return nil, result.Error
		}
	}

	return &dailyRecord, nil
}

func (r *ConsumeProductRepository) Create(product *entity.ConsumedProduct) error {
	if err := r.db.Create(product).Error; err != nil {
		return err
	}
	return nil
}

func (r *ConsumeProductRepository) GetByID(id uint) (*entity.ConsumedProduct, error) {
	var product entity.ConsumedProduct
	result := r.db.First(&product, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}

func (r *ConsumeProductRepository) DeleteByID(id uint) error {
	result := r.db.Delete(&entity.ConsumedProduct{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("product not found")
	}
	return nil
}

func (r *ConsumeProductRepository) CleanupDailyRecord(dailyID uint) error {
	var count int64
	r.db.Model(&entity.ConsumedProduct{}).Where("daily_products_consumed_id = ?", dailyID).Count(&count)

	if count == 0 {
		// No more consumed products, delete the DailyProductsConsumed entry
		if err := r.db.Delete(&entity.DailyProductsConsumed{}, dailyID).Error; err != nil {
			return err
		}
	}
	return nil
}
