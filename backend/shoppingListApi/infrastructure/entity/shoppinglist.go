package entity

type ShoppingList struct {
	ID          int       `json:"ID" gorm:"primaryKey"`
	Name        string    `json:"Name"`
	Description string    `json:"Description"`
	Date        string    `json:"Date"`
	Products    []Product `json:"Products" gorm:"foreignKey:ShoppingListID;constraint:OnDelete:CASCADE;"`
}

type Product struct {
	ID             int    `json:"ID" gorm:"primaryKey"`
	Name           string `json:"Name"`
	Quantity       string `json:"Quantity"`
	ShoppingListID int    `json:"ShoppingListID"`
}
