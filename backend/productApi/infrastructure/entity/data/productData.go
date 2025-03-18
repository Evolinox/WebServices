package data

import "productApi/infrastructure/entity"

var Products = []*entity.Product{
	{ID: 1, Name: "Skyr", Brand: "K-Classic", CaloriesPer100Grams: 64, ProteinsInGrams: 11, FatsInGrams: 0.2, CarbsInGrams: 4},
	{ID: 2, Name: "Protein Riegel", Brand: "ESN", CaloriesPer100Grams: 395, ProteinsInGrams: 28, FatsInGrams: 18, CarbsInGrams: 35},
	{ID: 3, Name: "Kartoffeln", Brand: "K-Classic", CaloriesPer100Grams: 76, ProteinsInGrams: 2, FatsInGrams: 0.5, CarbsInGrams: 16},
	{ID: 4, Name: "Skyr", Brand: "Milbona", CaloriesPer100Grams: 62, ProteinsInGrams: 11, FatsInGrams: 0.2, CarbsInGrams: 4},
	{ID: 5, Name: "Whey Protein", Brand: "ESN", CaloriesPer100Grams: 377, ProteinsInGrams: 72, FatsInGrams: 4.8, CarbsInGrams: 8.8},
	{ID: 6, Name: "Hähnchenbrustfilet", Brand: "Landjunker", CaloriesPer100Grams: 107, ProteinsInGrams: 23, FatsInGrams: 2.6, CarbsInGrams: 0.5},
	{ID: 7, Name: "Weizenbrot", Brand: "Bäcker", CaloriesPer100Grams: 265, ProteinsInGrams: 9, FatsInGrams: 3.2, CarbsInGrams: 49},
	{ID: 8, Name: "Bier", Brand: "Wulle", CaloriesPer100Grams: 43, ProteinsInGrams: 0.5, FatsInGrams: 0, CarbsInGrams: 3.6},
	{ID: 9, Name: "Frischkäse", Brand: "Philadelphia", CaloriesPer100Grams: 250, ProteinsInGrams: 5.5, FatsInGrams: 24, CarbsInGrams: 4},
	{ID: 10, Name: "Gouda", Brand: "G&G", CaloriesPer100Grams: 356, ProteinsInGrams: 25, FatsInGrams: 29, CarbsInGrams: 2},
	{ID: 11, Name: "Gurke", Brand: "Bauern", CaloriesPer100Grams: 12, ProteinsInGrams: 0.6, FatsInGrams: 0.1, CarbsInGrams: 2.2},
}
