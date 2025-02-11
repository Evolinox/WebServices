package mapper

import (
	"tracker/infrastructure/dto/model"
	"tracker/infrastructure/entity"
)

func CalculateSettingsData(settingsData model.SettingsDTO) entity.Settings {
	proteinsPercentage := 35
	fatsPercentage := 25
	carbsPercentage := 40
	proteinsInGrams := (settingsData.PlannedCalories * proteinsPercentage) / 100
	fatsInGrams := (settingsData.PlannedCalories * fatsPercentage) / 100
	carbsInGrams := (settingsData.PlannedCalories * carbsPercentage) / 100
	return entity.Settings{
		PlannedCalories: settingsData.PlannedCalories,
		ProteinsInGrams: proteinsInGrams,
		FatsInGrams:     fatsInGrams,
		CarbsInGrams:    carbsInGrams,
	}

}
