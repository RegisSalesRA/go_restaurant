package main

import (
	"log"
	"restaurante/internal/config"
	"restaurante/internal/database"
	"restaurante/internal/handle"     // Import direto para evitar confusão
	"restaurante/internal/repository" // Você vai precisar importar o repository aqui

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Carrega Configurações
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}

	// 2. Conecta ao Banco de Dados
	pool, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer pool.Close()

	// --- A PARTE QUE FALTAVA ---
	
	// 3. Instancia o Repository (Passa o pool de conexão)
	cashRepo := repository.NewCashDrawerRepository(pool)

	// 4. Instancia o Handler (Injeta o repository que acabamos de criar)
	cashHandler := &handlers.CashDrawerHandler{Repo: cashRepo}

	// ---------------------------

	router := gin.Default()
	router.SetTrustedProxies(nil)

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message":  "Restaurante API is running well!",
			"status":   "success",
			"database": "connected",
		})
	})

	// 5. Agora sim, você usa a variável cashHandler que foi iniciada acima
	router.POST("/caixa/abrir", cashHandler.AbrirCaixaHandler)

	log.Printf("Server starting on port %s", cfg.Port)
	router.Run(":" + cfg.Port)
}