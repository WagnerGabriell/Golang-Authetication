package usecase

import (
	Dto "GolangAuthetication/internal/DTO"
	"GolangAuthetication/internal/entity"
	"GolangAuthetication/internal/infra/repository"
)

type CreateProductUseCase struct {
	ProductRepositoryMysql repository.IProductRepositoryMysql
}

func NewCreateProductUseCase(ProductRepositoryMysql repository.IProductRepositoryMysql) *CreateProductUseCase {
	return &CreateProductUseCase{ProductRepositoryMysql: ProductRepositoryMysql}
}

func (u *CreateProductUseCase) Execute(ProductDTOInput Dto.ProductDTOInput, headerId string) (*Dto.ProductDTOOutput, error) {
	product := entity.NewProduct(ProductDTOInput.Name, ProductDTOInput.Price, headerId)
	err := u.ProductRepositoryMysql.CreateProduct(product)
	if err != nil {
		return &Dto.ProductDTOOutput{}, err
	}

	return &Dto.ProductDTOOutput{
		Id:     product.Id,
		Name:   product.Name,
		Price:  product.Price,
		IdUser: product.UserId,
	}, nil
}
