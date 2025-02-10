package mapper

import (
	"tracker/infrastructure/dto/model"
	"tracker/infrastructure/entity"
)

func CalculateSettingsData(settingsData model.SettingsDTO) entity.Settings {
	proteinsInGrams := (settingsData.PlannedCalories * settingsData.ProteinsPercentage) / 100
	fatsInGrams := (settingsData.PlannedCalories * settingsData.FatsPercentage) / 100
	carbsInGrams := (settingsData.PlannedCalories * settingsData.CarbsPercentage) / 100
	return entity.Settings{
		PlannedCalories: settingsData.PlannedCalories,
		ProteinsInGrams: proteinsInGrams,
		FatsInGrams:     fatsInGrams,
		CarbsInGrams:    carbsInGrams,
	}

}
