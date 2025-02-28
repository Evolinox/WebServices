package repositories

type ShopListAPIRepository struct {
	baseURL string
}

func NewShopListAPIRepository(baseURL string) *ShopListAPIRepository {
	return &ShopListAPIRepository{baseURL: baseURL}
}
