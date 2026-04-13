package repository

import (
	"context"
	"restaurante/internal/models"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

/*
Responsabilidade:
- Acessar o banco de dados
- Executar queries SQL
- Retornar dados já estruturados

Esse arquivo NÃO deve conter:
- lógica de negócio
- validações complexas
- regras de API

----------------------------------------
*/

// Struct que representa o repositório
// Guarda a conexão com o banco (pool)
type CashDrawerRepository struct {
	pool *pgxpool.Pool
}

// Construtor do repository (Dependency Injection)
// Recebe o pool do banco e retorna uma instância pronta
func NewCashDrawerRepository(p *pgxpool.Pool) *CashDrawerRepository {
	return &CashDrawerRepository{pool: p}
}

/*
========================================
ABRIR CAIXA
========================================

O que faz:
- Insere um novo registro no banco
- Representa a abertura do caixa no dia

Retorno:
- Struct DailyCashDrawer preenchida
- error caso algo falhe
*/
func (r *CashDrawerRepository) AbrirCaixa(ctx context.Context, valorInicial int) (*models.DailyCashDrawer, error) {

	// Query SQL para inserir o caixa
	// $1 é um parâmetro seguro (evita SQL Injection)
	// RETURNING retorna os dados inseridos (evita SELECT depois)
	query := `
		INSERT INTO daily_cash_drawer (initial_value, status)
		VALUES ($1, 'aberto')
		RETURNING id, opened_at, status, initial_value
	`

	// Struct que vai receber os dados do banco
	var drawer models.DailyCashDrawer

	// Executa a query e lê o resultado
	// QueryRow → espera apenas uma linha
	err := r.pool.QueryRow(ctx, query, valorInicial).Scan(
		&drawer.ID,           // ID gerado pelo banco
		&drawer.OpenedAt,     // Data de abertura
		&drawer.Status,       // Status (aberto)
		&drawer.InitialValue, // Valor inicial
	)

	// Tratamento de erro
	// %w mantém o erro original (boa prática)
	if err != nil {
		return nil, fmt.Errorf("falha ao abrir caixa: %w", err)
	}

	// Retorna o ponteiro da struct preenchida
	return &drawer, nil
}

/*
========================================
RESUMO
========================================

- Repository → camada de acesso ao banco
- Usa pgxpool (pool de conexões → performance)
- Usa context (controle de execução)
- Usa QueryRow + Scan (mapeamento direto)
- Usa RETURNING (evita queries extras)

Fluxo:

Service → Repository → Banco
*/