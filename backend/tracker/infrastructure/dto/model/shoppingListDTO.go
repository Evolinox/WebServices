package model

type ShoppingListDTO struct {
	ID          int                    `json:"ID"`
	Name        string                 `json:"Name"`
	Description string                 `json:"Description"`
	Date        string                 `json:"Date"`
	Products    []ShoppingListEntryDTO `json:"Products"`
}

type ShoppingListEntryDTO struct {
	ID             int    `json:"ID"`
	Name           string `json:"Name"`
	Quantity       string `json:"Quantity"`
	ShoppingListID int    `json:"ShoppingListID"`
}
