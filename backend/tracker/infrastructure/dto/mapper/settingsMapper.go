package mapper

import (
	"math"
	"tracker/infrastructure/dto/model"
	"tracker/infrastructure/entity"
)

func CalculateSettingsData(settingsData model.SettingsDTO) entity.Settings {
	proteinsPercentage := 0.35
	fatsPercentage := 0.25
	carbsPercentage := 0.40

	// multiply with percentage and divide through the factor gram to calorie
	proteinsInGrams := (float64(settingsData.PlannedCalories) * proteinsPercentage) / 4
	fatsInGrams := (float64(settingsData.PlannedCalories) * fatsPercentage) / 9
	carbsInGrams := (float64(settingsData.PlannedCalories) * carbsPercentage) / 4

	proteinsInGramsInt := int(math.Floor(proteinsInGrams))
	fatsInGramsInt := int(math.Floor(fatsInGrams))
	carbsInGramsInt := int(math.Floor(carbsInGrams))

	return entity.Settings{
		PlannedCalories: settingsData.PlannedCalories,
		ProteinsInGrams: proteinsInGramsInt,
		FatsInGrams:     fatsInGramsInt,
		CarbsInGrams:    carbsInGramsInt,
		Gender:          settingsData.Gender,
		WeightInKg:      settingsData.WeightInKg,
	}

}
