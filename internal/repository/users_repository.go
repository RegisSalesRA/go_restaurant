package repository

import (
	"context"
	"fmt"
	"restaurante/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UsersRepository struct {
	pool *pgxpool.Pool
}

func NewUsersRepository(p *pgxpool.Pool) *UsersRepository {
	return &UsersRepository{pool: p}
}
func (r *UsersRepository) CreateUser(ctx context.Context, user models.Users) error {
	query := `
		INSERT INTO users (firstName, lastName, email, password)
		VALUES ($1, $2, $3, $4)
	`
	_, err := r.pool.Exec(ctx, query,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password,
	)

	if err != nil {
		return fmt.Errorf("falha ao criar usuário: %w", err)
	}

	return nil
}
func (r *UsersRepository) GetUserByEmail(ctx context.Context, email string) (*models.Users, error) {
	query := `
		SELECT id, firstName, lastName, email, password, createdAt 
		FROM users 
		WHERE email = $1
	`
	var u models.Users
	err := r.pool.QueryRow(ctx, query, email).Scan(
		&u.ID,
		&u.FirstName,
		&u.LastName,
		&u.Email,
		&u.Password,
		&u.CreatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("usuário não encontrado: %w", err)
	}

	return &u, nil
}
func (r *UsersRepository) GetUserByID(ctx context.Context, id int) (*models.Users, error) {
	query := `
		SELECT id, firstName, lastName, email, password, createdAt 
		FROM users 
		WHERE id = $1
	`
	var u models.Users
	err := r.pool.QueryRow(ctx, query, id).Scan(
		&u.ID,
		&u.FirstName,
		&u.LastName,
		&u.Email,
		&u.Password,
		&u.CreatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("usuário não encontrado: %w", err)
	}

	return &u, nil
}
func (r *UsersRepository) UsersList(ctx context.Context) ([]models.Users, error) {
	query := `SELECT id, firstName, lastName, email, createdAt FROM users`

	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("erro ao listar usuários: %w", err)
	}
	defer rows.Close()

	var users []models.Users
	for rows.Next() {
		var u models.Users
		err := rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}
