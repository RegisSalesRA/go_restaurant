package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"restaurante/internal/handler"
	"restaurante/internal/repository"
)

func RegisterRoutes(r *gin.Engine, pool *pgxpool.Pool) {
    // 1. Instancia os Repositories (Camada de Dados)
    cashRepo := repository.NewCashDrawerRepository(pool)
    categoriesRepo := repository.NewCategoriesRepository(pool)

    // 2. Instancia os Handlers (Camada de Controle)
    // Ajustei de 'handlers' para 'handler' para bater com seu import
    cashH := &handler.CashDrawerHandler{Repo: cashRepo}
    categoriesHandle := &handler.CategoriesHandler{Repo: categoriesRepo}

    // 3. Define as Rotas
    api := r.Group("/api/v1")
    {
        // Rotas de Caixa
        api.POST("/caixa/abrir", cashH.AbrirCaixaHandler)
        
        // Rotas de Categorias
        api.POST("/categories", categoriesHandle.SaveCategory)
        api.GET("/categories", categoriesHandle.CategoriesList)
        api.GET("/categories/search", categoriesHandle.SearchCategoriesHandler)
        api.GET("/categories/:id", categoriesHandle.GetCategoryByID)
        api.PUT("/categories/:id", categoriesHandle.UpdateCategory)
        api.DELETE("/categories/:id", categoriesHandle.DeleteCategory)
    }
}