package entity

type Settings struct {
	PlannedCalories int `json:"PlannedCalories"`
	ProteinsInGrams int `json:"ProteinsInGrams"`
	FatsInGrams     int `json:"FatsInGrams"`
	CarbsInGrams    int `json:"CarbsInGrams"`
}
