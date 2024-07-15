package entity

import "github.com/google/uuid"

type Product struct {
	Id     string
	Name   string
	Price  float32
	UserId string
}

func NewProduct(name string, price float32, userId string) *Product {
	return &Product{
		Id:     uuid.New().String(),
		Name:   name,
		Price:  price,
		UserId: userId,
	}
}
