package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"restaurante/internal/routes"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestCategoriesAPI(t *testing.T) {
	// 1. Configura o Gin em modo de teste para não poluir o terminal com logs
	gin.SetMode(gin.TestMode)
	engine := gin.Default()

	// 2. Registra as rotas usando o pool de teste inicializado no main_test.go
	routes.RegisterRoutes(engine, testPool)

	t.Run("GET /api/v1/categories - Deve retornar status 200", func(t *testing.T) {
		// 3. Cria uma requisição HTTP simulada
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/categories", nil)
		
		// 4. Cria um "Gravador" de resposta para capturar o que o servidor enviar
		w := httptest.NewRecorder()

		// 5. Executa a requisição no motor do Gin
		engine.ServeHTTP(w, req)

		// 6. Asserções (Validações)
		if w.Code != http.StatusOK {
			t.Errorf("Esperava status 200, recebeu %d. Response: %s", w.Code, w.Body.String())
		}

		// 7. Opcional: Validar se o corpo da resposta é um JSON válido (mesmo que vazio [])
		var resp []interface{}
		if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
			t.Errorf("O corpo da resposta não é um JSON válido: %v", err)
		}
	})

	t.Run("POST /api/v1/categories - Deve validar nome obrigatório", func(t *testing.T) {
		// Enviando um JSON vazio para forçar o erro de validação (binding:"required")
		req, _ := http.NewRequest(http.MethodPost, "/api/v1/categories", nil)
		w := httptest.NewRecorder()

		engine.ServeHTTP(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("Esperava 400 para input vazio, recebeu %d", w.Code)
		}
	})
}