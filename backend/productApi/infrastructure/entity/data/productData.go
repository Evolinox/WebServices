package data

import "productApi/infrastructure/entity"

var Products = []*entity.Product{
	{ID: 1, Name: "Skyr", Brand: "K-Classic", CaloriesPer100Grams: 64, ProteinsInGrams: 11, FatsInGrams: 0.2, CarbsInGrams: 4},
	{ID: 2, Name: "Protein Riegel", Brand: "ESN", CaloriesPer100Grams: 395, ProteinsInGrams: 28, FatsInGrams: 18, CarbsInGrams: 35},
	{ID: 3, Name: "Kartoffeln", Brand: "K-Classic", CaloriesPer100Grams: 76, ProteinsInGrams: 2, FatsInGrams: 0.5, CarbsInGrams: 16},
	{ID: 4, Name: "Skyr", Brand: "Milbona", CaloriesPer100Grams: 62, ProteinsInGrams: 11, FatsInGrams: 0.2, CarbsInGrams: 4},
	{ID: 5, Name: "Whey Protein", Brand: "ESN", CaloriesPer100Grams: 377, ProteinsInGrams: 72, FatsInGrams: 4.8, CarbsInGrams: 8.8},
	{ID: 6, Name: "Hähnchenbrustfilet", Brand: "Landjunker", CaloriesPer100Grams: 107, ProteinsInGrams: 23, FatsInGrams: 2.6, CarbsInGrams: 0.5},
}
