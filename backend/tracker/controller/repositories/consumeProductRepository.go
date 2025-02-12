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

func (r *ConsumeProductRepository) ConsumeProduct(consumedProduct *entity.ConsumedProduct) error {
	return r.db.Create(consumedProduct).Error
}

func (r *ConsumeProductRepository) DeleteConsumedProduct(productID uint) error {
	return r.db.Where("id = ?", productID).Delete(&entity.ConsumedProduct{}).Error
}

func (r *ConsumeProductRepository) GetByID(productID uint) (*entity.ConsumedProduct, error) {
	var consumedProduct entity.ConsumedProduct
	result := r.db.Where("id = ?", productID).First(&consumedProduct)
	if result.Error != nil {
		return nil, result.Error
	}
	return &consumedProduct, nil
}

func (r *ConsumeProductRepository) GetOrCreateDailyRecord(date string) (uint, error) {
	var record entity.DailyProductsConsumed
	result := r.db.Where("date = ?", date).First(&record)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			newRecord := entity.DailyProductsConsumed{
				Date: date,
			}
			if err := r.db.Create(&newRecord).Error; err != nil {
				return 0, err
			}
			return newRecord.ID, nil
		}
		return 0, result.Error
	}

	return record.ID, nil
}
