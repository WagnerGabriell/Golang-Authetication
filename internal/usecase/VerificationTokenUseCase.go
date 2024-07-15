package usecase

import (
	"fmt"

	"github.com/golang-jwt/jwt"
)

type VerificarTokenUseCase struct{}

func NewVerificarTokenUseCase() *VerificarTokenUseCase {
	return &VerificarTokenUseCase{}
}

func (u *VerificarTokenUseCase) Execute(tokenString string, Key []byte) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verifica o método de assinatura para garantir que corresponde ao nosso método esperado
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("token inválido")
		}
		return Key, nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", fmt.Errorf("invalid token or claims")
	}
	jsonData, ok := claims["Id"].(string)
	if !ok {
		return "", fmt.Errorf("token inválido")
	}
	return jsonData, nil
}
