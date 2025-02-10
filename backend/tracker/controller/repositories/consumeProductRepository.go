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

func (r *ConsumeProductRepository) Create(product *entity.ConsumedProduct) error {
	if err := r.db.Create(product).Error; err != nil {
		return err
	}
	return nil
}

func (r *ConsumeProductRepository) DeleteByID(id string) error {
	result := r.db.Delete(&entity.ConsumedProduct{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("product not found")
	}
	return nil
}
