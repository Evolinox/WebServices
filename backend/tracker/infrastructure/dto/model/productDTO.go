package model

type ProductDTO struct {
	ID                  string  `json:"id"`
	Name                string  `json:"name"`
	Brand               string  `json:"brand"`
	CaloriesPer100Grams int     `json:"caloriesPer100Grams"`
	ProteinsInGrams     float64 `json:"proteinsInGrams"`
	FatsInGrams         float64 `json:"fatsInGrams"`
	CarbsInGrams        float64 `json:"carbsInGrams"`
}
