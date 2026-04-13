package tests

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

// testPool exportado para o pacote de teste
var testPool *pgxpool.Pool

func TestMain(m *testing.M) {
    ctx := context.Background()

    // Tenta carregar da raiz se estiver rodando da pasta tests ou de uma subpasta
    // Tenta primeiro um nível acima, depois dois
    _ = godotenv.Load("../.env")     // Caso esteja em /tests
    _ = godotenv.Load("../../.env")  // Caso esteja em /cmd/tests

    connStr := os.Getenv("DATABASE_URL")
    if connStr == "" {
        // Se o .env falhar, definimos um fallback manual para não usar o padrão do sistema
        connStr = "SEU BANCO AQUI"
    }

    fmt.Printf("Conectando ao banco: %s\n", connStr)

    var err error
    testPool, err = pgxpool.New(ctx, connStr)
    if err != nil {
        fmt.Printf("Erro ao configurar pool: %v\n", err)
        os.Exit(1)
    }

    if err := testPool.Ping(ctx); err != nil {
        fmt.Printf("Falha na autenticação ou banco inacessível: %v\n", err)
        os.Exit(1)
    }

    os.Exit(m.Run())
}