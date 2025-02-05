package entity

type ProductDTO struct {
	ID                  string  `json:"id"`
	Name                string  `json:"name"`
	Brand               string  `json:"brand"`
	PortionSizeInGrams  int     `json:"portionSizeInGrams"`
	CaloriesPer100Grams int     `json:"caloriesPer100Grams"`
	ProteinsInGrams     float64 `json:"proteinsInGrams"`
	FatsInGrams         float64 `json:"fatsInGrams"`
	CarbsInGrams        float64 `json:"carbsInGrams"`
}
