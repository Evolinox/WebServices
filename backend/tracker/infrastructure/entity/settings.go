package entity

type Settings struct {
	PlannedCalories int     `json:"PlannedCalories"`
	ProteinsInGrams int     `json:"ProteinsInGrams"`
	FatsInGrams     int     `json:"FatsInGrams"`
	CarbsInGrams    int     `json:"CarbsInGrams"`
	Gender          string  `json:"Gender" gorm:"type:enum('männlich', 'weiblich', 'divers')"`
	WeightInKg      float64 `json:"WeightInKg"`
}
