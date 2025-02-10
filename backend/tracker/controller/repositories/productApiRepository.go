package repositories

import (
	"encoding/json"
	"fmt"
	"net/http"
	"tracker/infrastructure/dto/model"
)

type ProductAPIRepository struct {
	baseURL string
}

func NewProductAPIRepository(baseURL string) *ProductAPIRepository {
	return &ProductAPIRepository{baseURL: baseURL}
}

func (r *ProductAPIRepository) GetProducts() ([]model.ProductDTO, error) {
	url := fmt.Sprintf("%s/products", r.baseURL)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch products: %s", resp.Status)
	}

	var products []model.ProductDTO
	if err := json.NewDecoder(resp.Body).Decode(&products); err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductAPIRepository) GetProductByID(id string) (*model.ProductDTO, error) {
	url := fmt.Sprintf("%s/products/%s", r.baseURL, id)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, nil
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch product: %s", resp.Status)
	}

	var product model.ProductDTO
	if err := json.NewDecoder(resp.Body).Decode(&product); err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *ProductAPIRepository) GetProductByName(name string) ([]model.ProductDTO, error) {
	url := fmt.Sprintf("%s/products/name/%s", r.baseURL, name)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, nil
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch products by name: %s", resp.Status)
	}

	var products []model.ProductDTO
	if err := json.NewDecoder(resp.Body).Decode(&products); err != nil {
		return nil, err
	}
	return products, nil
}
