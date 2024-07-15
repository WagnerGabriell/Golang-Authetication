package usecase

import (
	Dto "GolangAuthetication/internal/DTO"
	"GolangAuthetication/internal/entity"
	"GolangAuthetication/internal/infra/repository"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type TokenGeneratorUseCase struct {
	Repository repository.IuserRepository
}

func NewTokenGenerator(repository repository.IuserRepository) *TokenGeneratorUseCase {
	return &TokenGeneratorUseCase{Repository: repository}
}

func (u *TokenGeneratorUseCase) Execute(UserInputDTO Dto.UserDTOInput, Key []byte) (string, error) {
	User := entity.NewUser(UserInputDTO.Name, UserInputDTO.Email, UserInputDTO.Password)
	UserSelected, err := u.Repository.GetPerEmail(User)
	if err != nil {
		return "", err
	}
	if UserInputDTO.Password != UserSelected.Password {
		return "", fmt.Errorf("invalid password or email")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Id":  UserSelected.Id,
		"exp": time.Now().Add(time.Hour * 1).Unix(), // Expira em 1 hora
	})
	tokenString, err := token.SignedString(Key)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
