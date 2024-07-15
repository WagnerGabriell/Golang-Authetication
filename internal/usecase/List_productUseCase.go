package usecase

import (
	Dto "GolangAuthetication/internal/DTO"
	"GolangAuthetication/internal/infra/repository"
)

type ListProductUseCase struct {
	Repository repository.IProductRepositoryMysql
}

func NewListProductUseCase(repository repository.IProductRepositoryMysql) *ListProductUseCase {
	return &ListProductUseCase{Repository: repository}
}
func (u *ListProductUseCase) Execute() ([]*Dto.ProductDTOOutput, error) {
	var ListProductDTOOutput []*Dto.ProductDTOOutput
	Product, err := u.Repository.GetProduct()
	if err != nil {
		return nil, err
	}
	for _, product := range Product {
		ProductDTOOutput := &Dto.ProductDTOOutput{
			Id:     product.Id,
			Name:   product.Name,
			Price:  product.Price,
			IdUser: product.UserId,
		}
		ListProductDTOOutput = append(ListProductDTOOutput, ProductDTOOutput)
	}
	return ListProductDTOOutput, nil
}
