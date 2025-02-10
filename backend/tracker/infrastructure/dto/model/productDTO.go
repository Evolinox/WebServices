package model

type ProductDTO struct {
	ID                  uint    `json:"ID"`
	Name                string  `json:"Name"`
	Brand               string  `json:"Brand"`
	CaloriesPer100Grams int     `json:"CaloriesPer100Grams"`
	ProteinsInGrams     float64 `json:"ProteinsInGrams"`
	FatsInGrams         float64 `json:"FatsInGrams"`
	CarbsInGrams        float64 `json:"CarbsInGrams"`
}
