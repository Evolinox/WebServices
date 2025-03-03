package entity

type ShoppingList struct {
	Id          int       `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Date        string    `json:"date"`
	Products    []Product `json:"products" gorm:"foreignKey:ShoppingListID;constraint:OnDelete:CASCADE;"`
}

type Product struct {
	Id             int    `json:"id" gorm:"primaryKey"`
	Name           string `json:"name"`
	Quantity       string `json:"quantity"`
	ShoppingListID int    `json:"shopping_list_id"`
}
