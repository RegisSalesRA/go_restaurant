package handlers

import (
    "net/http"
    "restaurante/internal/repository" 
    "github.com/gin-gonic/gin"
)

// 1. Mudamos para string, pois Categoria tem nome, não valor inicial
type CreateCategoryInput struct {
    Name string `json:"name" binding:"required,min=3"`
}

type CategoriesHandler struct {
    Repo *repository.CategoriesRepository
}

func (h *CategoriesHandler) SaveCategory(c *gin.Context) {
    // 2. Use a struct de input correta
    var input CreateCategoryInput

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "O nome da categoria é obrigatório"})
        return
    }

    // 3. Passamos o input.Name (string) para o repository
    category, err := h.Repo.SaveCategory(c.Request.Context(), input.Name)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, category)
}