package repository

import (
	"context"
	"fmt"
	"restaurante/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CategoriesRepository struct {
	pool *pgxpool.Pool
}

func NewCategoriesRepository(p *pgxpool.Pool) *CategoriesRepository {
	return &CategoriesRepository{pool: p}
}


func (r *CategoriesRepository) SaveCategory(ctx context.Context, name string) (*models.Categories, error) {

    query := `
        INSERT INTO categories (name)
        VALUES ($1)
        RETURNING id, name, created_at
    `

    var category models.Categories

    err := r.pool.QueryRow(ctx, query, name).Scan(
        &category.ID,
        &category.Name,
        &category.CreatedAt,
    )

    if err != nil {
        return nil, fmt.Errorf("falha ao salvar category: %w", err)
    }

    return &category, nil
}