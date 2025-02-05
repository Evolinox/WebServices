package entity

type Product struct {
	ID                  string  `json:"ID" `
	Name                string  `json:"Name"`
	Brand               string  `json:"Brand"`
	PortionSizeInGrams  int     `json:"PortionSizeInGrams"`
	CaloriesPer100Grams int     `json:"CaloriesPer100Grams"`
	ProteinsInGrams     float64 `json:"ProteinsInGrams"`
	FatsInGrams         float64 `json:"FatsInGrams"`
	CarbsInGrams        float64 `json:"CarbsInGrams"`
}
