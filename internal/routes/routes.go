package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"restaurante/internal/handler"
	"restaurante/internal/repository"
)

func RegisterRoutes(r *gin.Engine, pool *pgxpool.Pool) {
	cashRepo := repository.NewCashDrawerRepository(pool)
	cashH := &handlers.CashDrawerHandler{Repo: cashRepo}

	api := r.Group("/api/v1")
	{
		api.POST("/caixa/abrir", cashH.AbrirCaixaHandler)
	}
}
