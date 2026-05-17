package repository

import (
	"context"
	"fmt"
	"restaurante/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductRepository struct {
	pool *pgxpool.Pool
}

func NewProductRepository(p *pgxpool.Pool) *ProductRepository {
	return &ProductRepository{pool: p}
}

func (r *ProductRepository) SaveProduct(ctx context.Context, name string, price int, stockQuantity int, categoryId int) (*models.Product, error) {

	query := `
        INSERT INTO products (name, price, stock_quantity, category_id)
        VALUES ($1, $2, $3, $4)
        RETURNING id, name, price, stock_quantity, category_id
    `
	var product models.Product

	err := r.pool.QueryRow(ctx, query, name, price, stockQuantity, categoryId ).Scan(
		&product.ID,
		&product.Name,
		&product.Price,
		&product.StockQuantity,
		&product.CategoryId,
	)

	if err != nil {
		return nil, fmt.Errorf("falha ao salvar product: %w", err)
	}

	return &product, nil
}

func (r *ProductRepository) ProductList(ctx context.Context) ([]models.Product, error) {

	query := `SELECT id, name, price, stock_quantity, category_id FROM products`

	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("erro ao consultar product: %w", err)
	}
	defer rows.Close()

	Product := []models.Product{}

	for rows.Next() {
		var c models.Product
		err := rows.Scan(&c.ID, &c.Name, &c.Price, &c.StockQuantity, &c.CategoryId)
		if err != nil {
			return nil, fmt.Errorf("erro ao escanear product: %w", err)
		}
		Product = append(Product, c)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("erro durante a iteração das products: %w", err)
	}

	return Product, nil
}

func (r *ProductRepository) UpdateProduct(ctx context.Context, id int, newName string, newPrice int, newStockQuantity int, newCategoryId int) error {
    query := `
        UPDATE products 
        SET name = $1, price = $2, stock_quantity = $3, category_id = $4
        WHERE id = $5
    `
    result, err := r.pool.Exec(ctx, query, newName, newPrice, newStockQuantity, newCategoryId, id)
    if err != nil {
        return fmt.Errorf("falha ao atualizar products: %w", err)
    }

    if result.RowsAffected() == 0 {
        return fmt.Errorf("product com ID %d não encontrada", id)
    }

    return nil
}

func (r *ProductRepository) DeleteProduct(ctx context.Context, id int) error {
	query := `DELETE FROM products WHERE id = $1`

	result, err := r.pool.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("falha ao deletar product: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("product com ID %d não encontrada", id)
	}

	return nil
}

func (r *ProductRepository) GetProductByID(ctx context.Context, id int) (*models.Product, error) {
	query := `SELECT id, name, price, stock_quantity FROM products WHERE id = $1`

	var product models.Product

	err := r.pool.QueryRow(ctx, query, id).Scan(
		&product.ID,
		&product.Name,
		&product.Price,
		&product.StockQuantity,
	)

	if err != nil {
		return nil, fmt.Errorf("product não encontrado: %w", err)
	}

	return &product, nil
}

func (r *ProductRepository) FilterProductByName(ctx context.Context, name string) ([]models.Product, error) {

	query := `
        SELECT *  
        FROM products 
        WHERE name ILIKE $1
    `

	searchTerm := "%" + name + "%"

	rows, err := r.pool.Query(ctx, query, searchTerm)
	if err != nil {
		return nil, fmt.Errorf("erro ao filtrar product: %w", err)
	}
	defer rows.Close()

	Product := []models.Product{}
	for rows.Next() {
		var c models.Product
		if err := rows.Scan(&c.ID, &c.Name, &c.Price, &c.StockQuantity); err != nil {
			return nil, err
		}
		Product = append(Product, c)
	}

	return Product, nil
}
