package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"restaurante/internal/repository"
	"restaurante/tests" // Seu helper que carrega o .env

	"github.com/gin-gonic/gin"
)

func TestCategoriesHandler(t *testing.T) {
	// Configura o Gin para não poluir o terminal com logs de debug
	gin.SetMode(gin.TestMode)

	// Inicializa conexão com banco de teste
	pool := tests.GetTestPool()
	defer pool.Close()

	// Injeta a dependência do repositório no handler
	repo := repository.NewCategoriesRepository(pool)
	h := &CategoriesHandler{Repo: repo}

	t.Run("POST /categories - Deve criar uma categoria com sucesso", func(t *testing.T) {
		r := gin.Default()
		r.POST("/categories", h.SaveCategory)

		input := CreateCategoryInput{Name: "Bebidas Alcoólicas"}
		jsonBody, _ := json.Marshal(input)

		req, _ := http.NewRequest(http.MethodPost, "/categories", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		if w.Code != http.StatusCreated {
			t.Errorf("Esperava 201, recebeu %d. Resposta: %s", w.Code, w.Body.String())
		}
	})

	t.Run("GET /categories - Deve listar categorias", func(t *testing.T) {
		r := gin.Default()
		r.GET("/categories", h.CategoriesList)

		req, _ := http.NewRequest(http.MethodGet, "/categories", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Esperava 200, recebeu %d", w.Code)
		}
	})

	t.Run("GET /categories/:id - Deve retornar 404 para ID inexistente", func(t *testing.T) {
		r := gin.Default()
		r.GET("/categories/:id", h.GetCategoryByID)

		// Usando um ID improvável (9999)
		req, _ := http.NewRequest(http.MethodGet, "/categories/9999", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		if w.Code != http.StatusNotFound {
			t.Errorf("Esperava 404 para categoria inexistente, recebeu %d", w.Code)
		}
	})

	t.Run("DELETE /categories/:id - Deve validar ID numérico", func(t *testing.T) {
		r := gin.Default()
		r.DELETE("/categories/:id", h.DeleteCategory)

		req, _ := http.NewRequest(http.MethodDelete, "/categories/abc", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("Esperava 400 para ID malformado, recebeu %d", w.Code)
		}
	})
}