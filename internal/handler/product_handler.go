package handler

import (
	"net/http"
	"restaurante/internal/repository"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CreateProductInput struct {
	Name          string `json:"name" binding:"required"`
	Price         int    `json:"price" binding:"required"`
	StockQuantity int    `json:"stock_quantity" binding:"required"`
	CategoryId int    `json:"category_id" binding:"required"`
}

type ProductsHandler struct {
	Repo *repository.ProductRepository
}

func (h *ProductsHandler) SaveProduct(c *gin.Context) {
	var input CreateProductInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "O nome da Product é obrigatório"})
		return
	}

	product, err := h.Repo.SaveProduct(c.Request.Context(), input.Name, input.Price, input.StockQuantity, input.CategoryId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, product)
}

func (h *ProductsHandler) ProductsList(c *gin.Context) {
	productsHandlers, err := h.Repo.ProductList(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, productsHandlers)
}

func (h *ProductsHandler) GetProductByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	product, err := h.Repo.GetProductByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product não encontrada"})
		return
	}
	c.JSON(http.StatusOK, product)
}

func (h *ProductsHandler) UpdateProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var input CreateProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	err = h.Repo.UpdateProduct(c.Request.Context(), id, input.Name, input.Price, input.StockQuantity, input.CategoryId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product atualizada com sucesso"})
}

func (h *ProductsHandler) DeleteProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = h.Repo.DeleteProduct(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product removida com sucesso"})
}

func (h *ProductsHandler) SearchProductsHandler(c *gin.Context) {
	queryName := c.Query("name")

	if queryName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "O parâmetro 'name' é obrigatório para a busca"})
		return
	}

	products, err := h.Repo.FilterProductByName(c.Request.Context(), queryName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao processar a busca no banco"})
		return
	}

	c.JSON(http.StatusOK, products)
}
