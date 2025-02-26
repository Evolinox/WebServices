package entity

type ShoppingList struct {
	Id          int    `json:"id" gorm:"primaryKey"`
	Date        string `json:"date"`
	BeginTime   string `json:"beginTime"`
	EndTime     string `json:"endTime"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
