package tests

import (
	"context"
	"os"
	"path/filepath"
	"runtime"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func GetTestPool() *pgxpool.Pool {
	
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)
	root := filepath.Join(basepath, "..")
	envPath := filepath.Join(root, ".env")
	
	err := godotenv.Load(envPath)

	connStr := os.Getenv("DATABASE_URL")
	

	pool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		panic("Falha ao conectar no banco: " + err.Error())
	}
	
	return pool
}