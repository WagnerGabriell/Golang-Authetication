package Dto

type ProductDTOInput struct {
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}
type ProductDTOOutput struct {
	Id     string
	Name   string
	Price  float32
	IdUser string
}
