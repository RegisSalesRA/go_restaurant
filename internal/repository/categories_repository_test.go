package repository

import (
	"context"
	"restaurante/tests" // Importa o pacote de testes da raiz
	"testing"
)

func TestCategoriesRepository(t *testing.T) {
	// 1. Pega uma conexão limpa usando o helper global
	pool := tests.GetTestPool()
	defer pool.Close() // Fecha após terminar os testes desta função

	repo := NewCategoriesRepository(pool)
	ctx := context.Background()

	t.Run("Deve salvar uma categoria com sucesso", func(t *testing.T) {
		name := "Massas"
		cat, err := repo.SaveCategory(ctx, name)

		if err != nil {
			t.Fatalf("Erro ao salvar: %v", err)
		}
		if cat.Name != name {
			t.Errorf("Nome esperado %s, recebeu %s", name, cat.Name)
		}
	})
}