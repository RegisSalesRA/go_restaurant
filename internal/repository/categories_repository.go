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

func (r *CategoriesRepository) CategoriesList(ctx context.Context) ([]models.Categories, error) {

	query := `SELECT id, name, created_at FROM categories`

	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("erro ao consultar categorias: %w", err)
	}
	defer rows.Close()

	categories := []models.Categories{}

	for rows.Next() {
		var c models.Categories
		err := rows.Scan(&c.ID, &c.Name, &c.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("erro ao escanear categoria: %w", err)
		}
		categories = append(categories, c)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("erro durante a iteração das categorias: %w", err)
	}

	return categories, nil
}


func (r *CategoriesRepository) UpdateCategory(ctx context.Context, id int, newName string) error {
    query := `
        UPDATE categories 
        SET name = $1 
        WHERE id = $2
    `
    result, err := r.pool.Exec(ctx, query, newName, id)
    if err != nil {
        return fmt.Errorf("falha ao atualizar categoria: %w", err)
    }

    if result.RowsAffected() == 0 {
        return fmt.Errorf("categoria com ID %d não encontrada", id)
    }

    return nil
}

func (r *CategoriesRepository) DeleteCategory(ctx context.Context, id int) error {
    query := `DELETE FROM categories WHERE id = $1`

    result, err := r.pool.Exec(ctx, query, id)
    if err != nil {
        return fmt.Errorf("falha ao deletar categoria: %w", err)
    }

    if result.RowsAffected() == 0 {
        return fmt.Errorf("categoria com ID %d não encontrada", id)
    }

    return nil
}

func (r *CategoriesRepository) GetCategoryByID(ctx context.Context, id int) (*models.Categories, error) {
    query := `SELECT id, name, created_at FROM categories WHERE id = $1`

    var category models.Categories

    err := r.pool.QueryRow(ctx, query, id).Scan(
        &category.ID,
        &category.Name,
        &category.CreatedAt,
    )

    if err != nil {
        return nil, fmt.Errorf("categoria não encontrada: %w", err)
    }

    return &category, nil
}


func (r *CategoriesRepository) FilterCategoriesByName(ctx context.Context, name string) ([]models.Categories, error) {

    query := `
        SELECT id, name, created_at 
        FROM categories 
        WHERE name ILIKE $1
    `

    searchTerm := "%" + name + "%"

    rows, err := r.pool.Query(ctx, query, searchTerm)
    if err != nil {
        return nil, fmt.Errorf("erro ao filtrar categorias: %w", err)
    }
    defer rows.Close()

    categories := []models.Categories{}
    for rows.Next() {
        var c models.Categories
        if err := rows.Scan(&c.ID, &c.Name, &c.CreatedAt); err != nil {
            return nil, err
        }
        categories = append(categories, c)
    }

    return categories, nil
}

