package model

type SettingsDTO struct {
	PlannedCalories int    `json:"PlannedCalories"`
	Gender          string `json:"Gender"`
	WeightInKg      int    `json:"WeightInKg"`
}
