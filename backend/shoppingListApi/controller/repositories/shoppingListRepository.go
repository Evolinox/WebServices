package repositories

import (
	"errors"
	"gorm.io/gorm"
	"shoppingListApi/infrastructure/entity"
)

type ShoppingListRepository struct {
	db *gorm.DB
}

func NewShoppingListRepository(db *gorm.DB) *ShoppingListRepository {
	return &ShoppingListRepository{db: db}
}

func (r *ShoppingListRepository) GetShoppingLists() ([]entity.ShoppingList, error) {
	var lists []entity.ShoppingList
	err := r.db.Preload("Products").Find(&lists).Error
	if err != nil {
		return nil, err
	}
	return lists, nil
}

func (r *ShoppingListRepository) GetShoppingListById(id int) (*entity.ShoppingList, error) {
	var list entity.ShoppingList
	err := r.db.Preload("Products").First(&list, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &list, nil
}

func (r *ShoppingListRepository) CreateShoppingList(list *entity.ShoppingList) error {
	return r.db.Create(list).Error
}

func (r *ShoppingListRepository) CreateShoppingListEntry(product *entity.Product) error {
	return r.db.Create(product).Error
}

func (r *ShoppingListRepository) DeleteShoppingList(id int) error {
	// Delete all associated products first
	if err := r.db.Where("shopping_list_id = ?", id).Delete(&entity.Product{}).Error; err != nil {
		return err
	}
	// Delete the shopping list
	return r.db.Delete(&entity.ShoppingList{}, id).Error
}

func (r *ShoppingListRepository) DeleteShoppingListEntry(productId int) error {
	return r.db.Delete(&entity.Product{}, productId).Error
}
