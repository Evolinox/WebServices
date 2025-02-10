package mapper

import (
	"tracker/infrastructure/dto/model"
	entity2 "tracker/infrastructure/entity"
)

func CalculateNutritionValues(productData model.ReceiveProductDTO) entity2.ConsumedProduct {
	weightFactor := float64(productData.Weight) / 100.0

	// Create a new ConsumedProduct
	consumedProduct := entity2.ConsumedProduct{
		ProductID:       productData.Product.ID,
		Date:            productData.Date,
		Category:        productData.Category,
		WeightInGrams:   productData.Weight,
		Calories:        int(float64(productData.Product.CaloriesPer100Grams) * weightFactor),
		ProteinsInGrams: productData.Product.ProteinsInGrams * weightFactor,
		FatsInGrams:     productData.Product.FatsInGrams * weightFactor,
		CarbsInGrams:    productData.Product.CarbsInGrams * weightFactor,
	}

	// Return the ConsumedProduct struct
	return consumedProduct
}
