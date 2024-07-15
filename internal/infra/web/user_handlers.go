package web

import (
	Dto "GolangAuthetication/internal/DTO"
	"GolangAuthetication/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandlers struct {
	CreateUserUsecase    *usecase.CreateUserUseCase
	ListUserUseCase      *usecase.ListUserUseCase
	LoginUserUseCase     *usecase.TokenGeneratorUseCase
	VerificaTokenUsecase *usecase.VerificarTokenUseCase
	GetUserUseCase       *usecase.GetUserUseCase
}

func NewUserHandlers(createUserUseCase *usecase.CreateUserUseCase, listUserUseCase *usecase.ListUserUseCase, tokenGeneratorUseCase *usecase.TokenGeneratorUseCase, verificaTokenUsecase *usecase.VerificarTokenUseCase, getUserUseCase *usecase.GetUserUseCase) *UserHandlers {
	return &UserHandlers{
		CreateUserUsecase:    createUserUseCase,
		ListUserUseCase:      listUserUseCase,
		LoginUserUseCase:     tokenGeneratorUseCase,
		VerificaTokenUsecase: verificaTokenUsecase,
		GetUserUseCase:       getUserUseCase,
	}
}
func (u *UserHandlers) CreateUserUseCase(c *gin.Context) {
	var userDTOInput Dto.UserDTOInput
	err := c.ShouldBindJSON(&userDTOInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	output, err := u.CreateUserUsecase.Execute(userDTOInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"user": output})
}

func (u *UserHandlers) ListUserUsecase(c *gin.Context) {
	output, err := u.ListUserUseCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"users": output})
}
func (u *UserHandlers) GetUserUsecase(c *gin.Context) {
	var userDTOInput Dto.UserDTOInput
	err := c.ShouldBindJSON(&userDTOInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	output, err := u.GetUserUseCase.Execute(userDTOInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"users": output})
}

func (u *UserHandlers) LoginUserUsecase(c *gin.Context) {
	var userDTOInput Dto.UserDTOInput
	err := c.ShouldBindJSON(&userDTOInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	var mySigningKey = []byte("minhaChaveSecreta")
	token, err := u.LoginUserUseCase.Execute(userDTOInput, mySigningKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}
func (u *UserHandlers) VerificarTokenUseCase(c *gin.Context) {
	jwtToken := c.GetHeader("Token")
	if jwtToken == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token invalid"})
		c.Abort()
	}
	var mySigningKey = []byte("minhaChaveSecreta")
	id, err := u.VerificaTokenUsecase.Execute(jwtToken, mySigningKey)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err})
		c.Abort()
	}
	c.Set("Id", id)
	c.Next()
}
