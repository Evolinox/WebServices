package mapper

import (
	"tracker/infrastructure/dto"
	"tracker/infrastructure/entity"
)

/*
	Flexibility for Future Changes
	- The mapper package is created to separate the conversion logic from the rest of the code.
	- This way, if the structure of the Product entity changes, only the conversion logic in the mapper package needs to be updated.
*/

func ToProductDTO(product entity.Product) dto.ProductDTO {
	return dto.ProductDTO{
		ID:                  product.ID,
		Name:                product.Name,
		Brand:               product.Brand,
		PortionSizeInGrams:  product.PortionSizeInGrams,
		CaloriesPer100Grams: product.CaloriesPer100Grams,
		ProteinsInGrams:     product.ProteinsInGrams,
		FatsInGrams:         product.FatsInGrams,
		CarbsInGrams:        product.CarbsInGrams,
	}
}

func ToProductEntity(dto dto.ProductDTO) entity.Product {
	return entity.Product{
		ID:                  dto.ID,
		Name:                dto.Name,
		Brand:               dto.Brand,
		PortionSizeInGrams:  dto.PortionSizeInGrams,
		CaloriesPer100Grams: dto.CaloriesPer100Grams,
		ProteinsInGrams:     dto.ProteinsInGrams,
		FatsInGrams:         dto.FatsInGrams,
		CarbsInGrams:        dto.CarbsInGrams,
	}
}
