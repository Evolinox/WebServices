package api

import "tracker/controller/repositories"

type DIC struct {
	productAPIRepo *repositories.ProductAPIRepository
}

func (d *DIC) GetProductAPIRepository() *repositories.ProductAPIRepository {
	if d.productAPIRepo == nil {
		d.productAPIRepo = repositories.NewProductAPIRepository("http://localhost:8081")
	}
	return d.productAPIRepo
}
