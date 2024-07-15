package repository

import "GolangAuthetication/internal/entity"

type IProductRepositoryMysql interface {
	GetProduct() ([]*entity.Product, error)
	CreateProduct(Product *entity.Product) error
}
