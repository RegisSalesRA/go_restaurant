package handler

import (
    "net/http"
    "restaurante/internal/repository" 
    "strconv" // Necessário para converter ID de string para int
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

func (h *CategoriesHandler) CategoriesList(c *gin.Context) {
    // A variável 'categories' agora será um slice []models.Categories
    categories, err := h.Repo.CategoriesList(c.Request.Context())
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
 
    c.JSON(http.StatusOK, categories)
}

func (h *CategoriesHandler) GetCategoryByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	category, err := h.Repo.GetCategoryByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Categoria não encontrada"})
		return
	}
	c.JSON(http.StatusOK, category)
}

func (h *CategoriesHandler) UpdateCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var input CreateCategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	err = h.Repo.UpdateCategory(c.Request.Context(), id, input.Name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Categoria atualizada com sucesso"})
}

func (h *CategoriesHandler) DeleteCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = h.Repo.DeleteCategory(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Categoria removida com sucesso"})
}


func (h *CategoriesHandler) SearchCategoriesHandler(c *gin.Context) { 
    queryName := c.Query("name")
 
    if queryName == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "O parâmetro 'name' é obrigatório para a busca"})
        return
    }
 
    categories, err := h.Repo.FilterCategoriesByName(c.Request.Context(), queryName)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao processar a busca no banco"})
        return
    }
 
    c.JSON(http.StatusOK, categories)
}