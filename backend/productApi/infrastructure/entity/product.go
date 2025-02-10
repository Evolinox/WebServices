package entity

type Product struct {
	ID                  uint    `json:"ID" gorm:"primaryKey"`
	Name                string  `json:"Name"`
	Brand               string  `json:"Brand"`
	CaloriesPer100Grams int     `json:"CaloriesPer100Grams"`
	ProteinsInGrams     float64 `json:"ProteinsInGrams"`
	FatsInGrams         float64 `json:"FatsInGrams"`
	CarbsInGrams        float64 `json:"CarbsInGrams"`
}
