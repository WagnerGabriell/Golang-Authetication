package usecase

import (
	Dto "GolangAuthetication/internal/DTO"
	"GolangAuthetication/internal/entity"
	"GolangAuthetication/internal/infra/repository"
)

type CreateUserUseCase struct {
	UserRepository repository.IuserRepository
}

func NewCreateUserUseCase(repository repository.IuserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{UserRepository: repository}
}

func (u *CreateUserUseCase) Execute(Input Dto.UserDTOInput) (*Dto.UserDTOOutput, error) {
	user := entity.NewUser(Input.Name, Input.Email, Input.Password)
	err := u.UserRepository.CreateUser(user)
	if err != nil {
		return &Dto.UserDTOOutput{}, err
	}
	return &Dto.UserDTOOutput{
		Id:       user.Id,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}, nil
}
