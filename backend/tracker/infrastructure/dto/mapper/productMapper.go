package mapper

import (
	"tracker/infrastructure/dto/model"
	entity2 "tracker/infrastructure/entity"
)

func CalculateNutritionValues(productData model.ReceiveProductDTO, dailyRecordID uint) *entity2.ConsumedProduct {
	weightFactor := float64(productData.Weight) / 100.0

	// Create the ConsumedProduct with the correct foreign key
	return &entity2.ConsumedProduct{
		DailyProductsConsumedID: dailyRecordID, // Assign the ID from the DailyProductsConsumed record
		ProductID:               productData.Product.ID,
		Category:                productData.Category,
		WeightInGrams:           productData.Weight,
		Calories:                int(float64(productData.Product.CaloriesPer100Grams) * weightFactor),
		ProteinsInGrams:         productData.Product.ProteinsInGrams * weightFactor,
		FatsInGrams:             productData.Product.FatsInGrams * weightFactor,
		CarbsInGrams:            productData.Product.CarbsInGrams * weightFactor,
	}
}
