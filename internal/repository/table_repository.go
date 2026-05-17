package repository

import (
	"context"
	"fmt"
	"restaurante/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TablesRepository struct {
	pool *pgxpool.Pool
}

func NewTablesRepository(p *pgxpool.Pool) *TablesRepository {
	return &TablesRepository{pool: p}
}

func (r *TablesRepository) SaveTables(ctx context.Context, number int, status string) (*models.Tables, error) {
	query := `
        INSERT INTO tables (number, status)
        VALUES ($1, $2)
        RETURNING id, number, status
    `
	var tables models.Tables

	err := r.pool.QueryRow(ctx, query, number, status).Scan(
		&tables.ID,
		&tables.Number,
		&tables.Status,
	)

	if err != nil {
		return nil, fmt.Errorf("falha ao salvar tables: %w", err)
	}

	return &tables, nil
}

func (r *TablesRepository) TablesList(ctx context.Context) ([]models.Tables, error) {

	query := `SELECT id, number, status FROM tables`

	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("erro ao consultar tables: %w", err)
	}
	defer rows.Close()

	tables := []models.Tables{}

	for rows.Next() {
		var c models.Tables
		err := rows.Scan(&c.ID, &c.Number, &c.Status)
		if err != nil {
			return nil, fmt.Errorf("erro ao escanear tables: %w", err)
		}
		tables = append(tables, c)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("erro durante a iteração das tables: %w", err)
	}

	return tables, nil
}

func (r *TablesRepository) UpdateTable(ctx context.Context, id int, status string) error {
	query := `
        UPDATE tables 
        SET status = $2  
        WHERE id = $1
    `
	result, err := r.pool.Exec(ctx, query, id, status)
	if err != nil {
		return fmt.Errorf("falha ao atualizar tables: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("tables com ID %d não encontrada", id)
	}

	return nil
}

func (r *TablesRepository) DeleteTable(ctx context.Context, id int) error {
	query := `DELETE FROM tables WHERE id = $1`

	result, err := r.pool.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("falha ao deletar tables: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("tables com ID %d não encontrada", id)
	}

	return nil
}

func (r *TablesRepository) GetTableByID(ctx context.Context, id int) (*models.Tables, error) {
	query := `SELECT id, number, status FROM tables WHERE id = $1`

	var tables models.Tables

	err := r.pool.QueryRow(ctx, query, id).Scan(
		&tables.ID,
		&tables.Number,
		&tables.Status,
	)

	if err != nil {
		return nil, fmt.Errorf("tables não encontrada: %w", err)
	}

	return &tables, nil
}

func (r *TablesRepository) FilterTablesByName(ctx context.Context, name string) ([]models.Tables, error) {

	query := `
        SELECT id, number, status
        FROM tables 
        WHERE name ILIKE $1
    `

	searchTerm := "%" + name + "%"

	rows, err := r.pool.Query(ctx, query, searchTerm)
	if err != nil {
		return nil, fmt.Errorf("erro ao filtrar tables: %w", err)
	}
	defer rows.Close()

	tables := []models.Tables{}
	for rows.Next() {
		var c models.Tables
		if err := rows.Scan(&c.ID, &c.Number, &c.Status); err != nil {
			return nil, err
		}
		tables = append(tables, c)
	}

	return tables, nil
}
