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

func (r *ShoppingListRepository) DeleteShoppingList(id string) error {
	res := r.db.Delete(&entity.ShoppingList{}, "id = ?", id)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("shopping list not found")
	}
	return nil
}

func (r *ShoppingListRepository) DeleteShoppingListEntry(id string, productId string) error {
	res := r.db.Delete(&entity.Product{}, "id = ? AND shopping_list_id = ?", productId, id)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("shopping list entry not found")
	}
	return nil
}
