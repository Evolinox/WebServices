package repositories

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"tracker/infrastructure/dto/model"
)

type ShopListAPIRepository struct {
	baseURL string
}

func NewShopListAPIRepository(baseURL string) *ShopListAPIRepository {
	return &ShopListAPIRepository{baseURL: baseURL}
}

func (r *ShopListAPIRepository) GetShoppingLists() ([]model.ShoppingListDTO, error) {
	url := fmt.Sprintf("%s/shoppinglist/", r.baseURL)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch shopping lists: %s", resp.Status)
	}

	var lists []model.ShoppingListDTO
	if err := json.NewDecoder(resp.Body).Decode(&lists); err != nil {
		return nil, err
	}
	return lists, nil
}

func (r *ShopListAPIRepository) GetShoppingListByID(id string) (*model.ShoppingListDTO, error) {
	url := fmt.Sprintf("%s/shoppinglist/%s", r.baseURL, id)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, nil
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch shopping list: %s", resp.Status)
	}

	var list model.ShoppingListDTO
	if err := json.NewDecoder(resp.Body).Decode(&list); err != nil {
		return nil, err
	}
	return &list, nil
}

func (r *ShopListAPIRepository) CreateShoppingList(list model.ShoppingListDTO) error {
	url := fmt.Sprintf("%s/shoppinglist/", r.baseURL)
	data, err := json.Marshal(list)
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("failed to create shopping list: %s", resp.Status)
	}
	return nil
}

func (r *ShopListAPIRepository) CreateShoppingListEntry(id string, entry model.ShoppingListEntryDTO) error {
	url := fmt.Sprintf("%s/shoppinglist/%s", r.baseURL, id)
	data, err := json.Marshal(entry)
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("failed to create shopping list entry: %s", resp.Status)
	}
	return nil
}

func (r *ShopListAPIRepository) DeleteShoppingList(id string) error {
	url := fmt.Sprintf("%s/shoppinglist/%s", r.baseURL, id)
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return fmt.Errorf("shopping list not found")
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to delete shopping list: %s", resp.Status)
	}
	return nil
}

func (r *ShopListAPIRepository) DeleteShoppingListEntry(id, entryId string) error {
	url := fmt.Sprintf("%s/shoppinglist/%s/products/%s", r.baseURL, id, entryId)
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return fmt.Errorf("shopping list entry not found")
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to delete shopping list entry: %s", resp.Status)
	}
	return nil
}
