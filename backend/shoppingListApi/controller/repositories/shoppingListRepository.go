package repositories

import (
	"gorm.io/gorm"
	"shoppingListApi/infrastructure/entity"
)

type ShoppingListRepository struct {
	db *gorm.DB
}

func NewShoppingListRepository(db *gorm.DB) *ShoppingListRepository {
	return &ShoppingListRepository{db: db}
}

func (repo *ShoppingListRepository) Create(entry *entity.ShoppingList) error {
	if err := repo.db.Create(entry).Error; err != nil {
		return err
	}
	return nil
}
