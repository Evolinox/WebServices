package model

type ReceiveProductDTO struct {
	Product  ProductDTO `json:"Product"`
	Weight   int        `json:"Weight"`
	Category string     `json:"Category"`
}
