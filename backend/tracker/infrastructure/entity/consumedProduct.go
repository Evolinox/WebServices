package entity

type ConsumedProduct struct {
	ID              uint    `json:"ID" gorm:"primaryKey;autoIncrement"`
	Date            string  `json:"Date"`
	ProductID       string  `json:"ProductID"`
	Category        string  `json:"Category" gorm:"type:enum('Frühstück', 'Mittagessen', 'Abendessen', 'Snack')"`
	WeightInGrams   int     `json:"WeightInGrams"`
	Calories        int     `json:"Calories"`
	ProteinsInGrams float64 `json:"ProteinsInGrams"`
	FatsInGrams     float64 `json:"FatsInGrams"`
	CarbsInGrams    float64 `json:"CarbsInGrams"`
}
