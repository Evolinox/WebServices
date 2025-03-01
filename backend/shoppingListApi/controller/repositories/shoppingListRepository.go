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

func (repo *ShoppingListRepository) Create(entry *entity.ShoppingList) error {
	if err := repo.db.Create(entry).Error; err != nil {
		return err
	}
	return nil
}

func (repo *ShoppingListRepository) GetAll() ([]entity.ShoppingList, error) {
	var shoppingLists []entity.ShoppingList
	if err := repo.db.Find(&shoppingLists).Error; err != nil {
		return nil, err
	}
	return shoppingLists, nil
}

func (repo *ShoppingListRepository) GetByDate(date string) ([]*entity.ShoppingList, error) {
	var list []*entity.ShoppingList
	if err := repo.db.Where("date = ?", date).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (repo *ShoppingListRepository) GetById(id string) (*entity.ShoppingList, error) {
	var list *entity.ShoppingList
	if err := repo.db.Where("id = ?", id).First(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (repo *ShoppingListRepository) GetByDateById(date string, id string) (*entity.ShoppingList, error) {
	var list *entity.ShoppingList
	if err := repo.db.Where("date = ? AND id = ?", date, id).First(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (repo *ShoppingListRepository) UpdateById(id string, updatedList *entity.ShoppingList) error {
	var list entity.ShoppingList
	if err := repo.db.Where("id = ?", id).First(&list).Error; err != nil {
		return err
	}
	return repo.db.Model(&entity.ShoppingList{}).Where("id = ?", id).Updates(updatedList).Error
}

func (repo *ShoppingListRepository) DeleteById(id string) error {
	result := repo.db.Where("id = ?", id).Delete(&entity.ShoppingList{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("shoppinglist not found")
	}
	return nil
}
