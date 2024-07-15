package main

import (
	"GolangAuthetication/internal/infra/repository"
	"GolangAuthetication/internal/infra/web"
	"GolangAuthetication/internal/usecase"
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:docker@tcp(localhost:3306)/authentication")
	router := gin.Default()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repositoryUser := repository.NewUserRepositoryMysql(db)
	repositoryProduct := repository.NewProductRepositoryMysql(db)

	createUserUsecase := usecase.NewCreateUserUseCase(repositoryUser)
	ListUserUseCase := usecase.NewListUserUseCase(repositoryUser)
	LoginUserUseCase := usecase.NewTokenGenerator(repositoryUser)
	GetUserUseCase := usecase.NewGetUserUseCase(repositoryUser)
	VerificarTokenUseCase := usecase.NewVerificarTokenUseCase()

	createProductUsecase := usecase.NewCreateProductUseCase(repositoryProduct)
	listProductUsecase := usecase.NewListProductUseCase(repositoryProduct)

	UserHandler := web.NewUserHandlers(createUserUsecase, ListUserUseCase, LoginUserUseCase, VerificarTokenUseCase, GetUserUseCase)
	ProductHandler := web.NewProductHandler(createProductUsecase, listProductUsecase)

	users := router.Group("/users")
	{
		users.POST("/create", UserHandler.CreateUserUseCase)
		users.GET("/list", UserHandler.ListUserUsecase)
		users.POST("/login", UserHandler.LoginUserUsecase)
		users.POST("/getEmail", UserHandler.GetUserUsecase)
	}
	product := router.Group("/product")
	{
		product.POST("/create", UserHandler.VerificarTokenUseCase, ProductHandler.CreateProduct)
		product.GET("/list", ProductHandler.ListProduct)
	}

	router.Run(":3000")
}
