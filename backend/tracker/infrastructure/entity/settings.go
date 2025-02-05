package entity

// ToDo Add helper Methods with calculator from Percentage to Grams and store grams in db
// Create DTO for this entity
type Settings struct {
	PlannedCalories    int `json:"PlannedCalories"`
	ProteinsPercentage int `json:"ProteinsPercentage"`
	FatsPercentage     int `json:"FatsPercentage"`
	CarbsPercentage    int `json:"CarbsPercentage"`
}
