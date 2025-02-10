package model

type ReceiveProductDTO struct {
	Product  ProductDTO `json:"Product"`
	Date     string     `json:"Date"`
	Weight   int        `json:"Weight"`
	Category string     `json:"Category"`
}
