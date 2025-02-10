package model

type SettingsDTO struct {
	PlannedCalories    int `json:"PlannedCalories"`
	ProteinsPercentage int `json:"ProteinsPercentage"`
	FatsPercentage     int `json:"FatsPercentage"`
	CarbsPercentage    int `json:"CarbsPercentage"`
}
