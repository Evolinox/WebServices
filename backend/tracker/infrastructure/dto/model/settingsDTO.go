package model

type SettingsDTO struct {
	PlannedCalories int     `json:"PlannedCalories"`
	Gender          string  `json:"Gender"`
	WeightInKg      float64 `json:"WeightInKg"`
}
