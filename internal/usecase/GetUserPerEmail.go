package usecase

import (
	Dto "GolangAuthetication/internal/DTO"
	"GolangAuthetication/internal/entity"
	"GolangAuthetication/internal/infra/repository"
)

type GetUserUseCase struct {
	UserRepository repository.IuserRepository
}

func NewGetUserUseCase(repository repository.IuserRepository) *GetUserUseCase {
	return &GetUserUseCase{UserRepository: repository}
}

func (u *GetUserUseCase) Execute(UserInputDTo Dto.UserDTOInput) (*Dto.UserDTOOutput, error) {
	User := entity.NewUser(UserInputDTo.Name, UserInputDTo.Email, UserInputDTo.Password)
	res, err := u.UserRepository.GetPerEmail(User)
	if err != nil {
		return &Dto.UserDTOOutput{}, err
	}

	return &Dto.UserDTOOutput{Id: res.Id, Name: res.Name, Email: res.Email, Password: res.Password}, nil
}
