package entity

type ConsumedProduct struct {
	ID                      uint    `json:"ID" gorm:"primaryKey"`
	DailyProductsConsumedID uint    `json:"DailyProductsConsumedID"`
	ProductID               uint    `json:"ProductID"`
	ProductName             string  `json:"ProductName"`
	Category                string  `json:"Category" gorm:"type:enum('Frühstück', 'Mittagessen', 'Abendessen', 'Snack')"`
	WeightInGrams           int     `json:"WeightInGrams"`
	Calories                int     `json:"Calories"`
	ProteinsInGrams         float64 `json:"ProteinsInGrams"`
	FatsInGrams             float64 `json:"FatsInGrams"`
	CarbsInGrams            float64 `json:"CarbsInGrams"`
}
