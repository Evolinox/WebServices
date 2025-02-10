package entity

type NutritionStatistics struct {
	ID               uint    `json:"ID" gorm:"primaryKey; Key:AutoIncrement"`
	Date             string  `json:"Date" gorm:"uniqueIndex"`
	ConsumedCalories int     `json:"ConsumedCalories"`
	ConsumedProteins float64 `json:"ConsumedProteins"`
	ConsumedFats     float64 `json:"ConsumedFats"`
	ConsumedCarbs    float64 `json:"ConsumedCarbs"`
}
