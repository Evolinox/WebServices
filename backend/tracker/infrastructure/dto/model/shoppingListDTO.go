package model

type ShoppingListDTO struct {
	Id          int                    `json:"id"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Date        string                 `json:"date"`
	Products    []ShoppingListEntryDTO `json:"products"`
}

type ShoppingListEntryDTO struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	Quantity       string `json:"quantity"`
	ShoppingListID int    `json:"shopping_list_id"`
}
