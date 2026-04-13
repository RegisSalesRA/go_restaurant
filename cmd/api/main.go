package main

import (
	"log"
	"restaurante/config"
	"restaurante/config/postgres"
	"restaurante/routes" 
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

	 
	router := gin.Default()
	router.SetTrustedProxies(nil)
    
    // Chamamos uma única função que resolve toda a bagunça
    routes.RegisterRoutes(router, pool)

	log.Printf("Server starting on port %s", cfg.Port)
	router.Run(":" + cfg.Port)
}