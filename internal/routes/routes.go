package routes 

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func RegisterRoutes(r *gin.Engine, pool *pgxpool.Pool) {
	cashRepo := repository.NewCashDrawerRepository(pool)
	prodRepo := repository.NewProductRepository(pool)

	cashH := &CashDrawerHandler{Repo: cashRepo}
	prodH := &ProductHandler{Repo: prodRepo}

	api := r.Group("/api/v1")
	{
		api.POST("/caixa/abrir", cashH.AbrirCaixaHandler)
		api.GET("/produtos", prodH.ListarProdutosHandler)
	}
}
