package usecase

import (
	Dto "GolangAuthetication/internal/DTO"
	"GolangAuthetication/internal/infra/repository"
)

type ListUserUseCase struct {
	UserRepository repository.IuserRepository
}

func NewListUserUseCase(repository repository.IuserRepository) *ListUserUseCase {
	return &ListUserUseCase{UserRepository: repository}
}

func (u *ListUserUseCase) Execute() ([]*Dto.UserDTOOutput, error) {
	users, err := u.UserRepository.GetUser()
	if err != nil {
		return nil, err
	}
	var ListUserDTOOutput []*Dto.UserDTOOutput
	for _, user := range users {
		ListUserDTOOutput = append(ListUserDTOOutput, &Dto.UserDTOOutput{
			Id:       user.Id,
			Name:     user.Name,
			Email:    user.Email,
			Password: user.Password,
		})
	}
	return ListUserDTOOutput, nil
}
