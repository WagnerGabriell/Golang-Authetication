package web

import (
	Dto "GolangAuthetication/internal/DTO"
	"GolangAuthetication/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	CreateProductUseCase *usecase.CreateProductUseCase
	ListProductUseCase   *usecase.ListProductUseCase
}

func NewProductHandler(createProductUseCase *usecase.CreateProductUseCase, listProductUseCase *usecase.ListProductUseCase) *ProductHandler {
	return &ProductHandler{
		CreateProductUseCase: createProductUseCase,
		ListProductUseCase:   listProductUseCase,
	}
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var productDTOInput Dto.ProductDTOInput
	err := c.ShouldBindJSON(&productDTOInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	headerId, exists := c.Get("Id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID não encontrado no contexto"})
		return
	}
	productDTOOutput, err := h.CreateProductUseCase.Execute(productDTOInput, headerId.(string))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"Product": productDTOOutput})
}

func (h *ProductHandler) ListProduct(c *gin.Context) {
	out, err := h.ListProductUseCase.Execute()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Products": out})
}
