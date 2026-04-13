package tests


import (
    "net/http"
    "net/http/httptest"
    "restaurante/routes"
    "testing"

    "github.com/gin-gonic/gin"
)

func TestCategoriesAPI(t *testing.T) {
    gin.SetMode(gin.TestMode)
    engine := gin.Default()

    // Chame a função diretamente para garantir que o pool existe
    pool := GetTestPool()
    defer pool.Close()

    // Registra as rotas usando o pool local
    routes.RegisterRoutes(engine, pool)

    t.Run("GET /api/v1/categories - Status 200", func(t *testing.T) {
        req, _ := http.NewRequest(http.MethodGet, "/api/v1/categories", nil)
        w := httptest.NewRecorder()
        engine.ServeHTTP(w, req)

        if w.Code != http.StatusOK {
            t.Errorf("Esperava 200, recebeu %d", w.Code)
        }
    })
    
    // ... outros sub-testes
}