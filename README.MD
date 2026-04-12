# 🍽️ Restaurant Ordering System - Database Schema

Este documento descreve o **schema do banco de dados** para um sistema
de gerenciamento de restaurante, incluindo:

-   Cardápio (categorias e produtos)
-   Mesas
-   Sessões de mesas (contas)
-   Pedidos
-   Pagamentos
-   Controle de caixa

O banco foi projetado para **PostgreSQL**.

------------------------------------------------------------------------

# 1. Menu: Categorias e Produtos

Responsável por armazenar os itens do cardápio.

## Categories

Categorias do cardápio (ex: Bebidas, Sobremesas).

``` sql
CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);
```

## Products

Itens disponíveis para venda.

-   `price` é armazenado em **centavos** para evitar erros de ponto
    flutuante.
-   `stock_quantity` controla o estoque.

``` sql
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    category_id INTEGER REFERENCES categories(id),
    name VARCHAR(100) NOT NULL,
    price INTEGER NOT NULL,
    stock_quantity INTEGER DEFAULT 0
);
```

------------------------------------------------------------------------

# 2. Infraestrutura: Mesas

Controle físico das mesas do restaurante.

``` sql
CREATE TABLE tables (
    id SERIAL PRIMARY KEY,
    number INTEGER UNIQUE NOT NULL,
    status VARCHAR(20) DEFAULT 'disponivel'
);
```

### Status possíveis

-   `disponivel`
-   `ocupada`
-   `suja`

------------------------------------------------------------------------

# 3. Operação: Sessões e Pedidos

## Table Sessions

Representa a **estadia de clientes em uma mesa** (a conta da mesa).

``` sql
CREATE TABLE table_sessions (
    id UUID PRIMARY KEY,
    table_id INTEGER REFERENCES tables(id),
    opened_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    closed_at TIMESTAMP WITH TIME ZONE,
    total_amount INTEGER DEFAULT 0
);
```

### Campos importantes

-   `opened_at` → horário de abertura da mesa
-   `closed_at` → horário de fechamento
-   `total_amount` → total acumulado da conta

------------------------------------------------------------------------

## Orders

Cada vez que um garçom envia um pedido para a cozinha.

``` sql
CREATE TABLE orders (
    id UUID PRIMARY KEY,
    session_id UUID REFERENCES table_sessions(id),
    status VARCHAR(20) DEFAULT 'pendente',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
```

### Status possíveis

-   `pendente`
-   `preparando`
-   `pronto`
-   `entregue`

------------------------------------------------------------------------

## Order Items

Itens dentro de um pedido.

``` sql
CREATE TABLE order_items (
    id SERIAL PRIMARY KEY,
    order_id UUID REFERENCES orders(id) ON DELETE CASCADE,
    product_id INTEGER REFERENCES products(id),
    quantity INTEGER NOT NULL,
    unit_price INTEGER NOT NULL
);
```

### Observação importante

`unit_price` captura o **preço no momento da venda**, preservando
histórico mesmo se o preço do produto mudar.

------------------------------------------------------------------------

# 4. Financeiro: Pagamentos e Caixa

## Payments

Registro dos pagamentos realizados.

``` sql
CREATE TABLE payments (
    id SERIAL PRIMARY KEY,
    session_id UUID REFERENCES table_sessions(id),
    amount INTEGER NOT NULL,
    payment_method VARCHAR(20),
    paid_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
```

### Métodos de pagamento

-   `dinheiro`
-   `cartao`
-   `pix`

------------------------------------------------------------------------

## Daily Cash Drawer

Controle do **caixa diário do restaurante**.

``` sql
CREATE TABLE daily_cash_drawer (
    id SERIAL PRIMARY KEY,
    opened_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    closed_at TIMESTAMP WITH TIME ZONE,
    initial_value INTEGER DEFAULT 0,    -- O "fundo de caixa" (dinheiro que começa para troco)
    total_recorded INTEGER DEFAULT 0,   -- O sistema soma aqui automaticamente a cada pedido pago
    final_counted INTEGER DEFAULT 0,    -- O valor que o usuário digita ao fechar o dia
    status VARCHAR(10) DEFAULT 'aberto' CHECK (status IN ('aberto', 'fechado'))
);
```

### Status

-   `aberto`
-   `fechado`

------------------------------------------------------------------------

# Dados para Teste Inicial

Inserção de alguns registros para testes.

``` sql
INSERT INTO categories (name) VALUES 
('Bebidas'),
('Pratos Principais');

INSERT INTO products (name, category_id, price, stock_quantity) VALUES 
('Coca-Cola 350ml', 1, 600, 50),
('Hambúrguer Gourmet', 2, 3500, 20);

INSERT INTO tables (number, status) VALUES 
(1, 'disponivel'),
(2, 'disponivel'),
(3, 'disponivel');
```

------------------------------------------------------------------------

# Estrutura Geral do Sistema

    categories
       │
       └── products
               │
               └── order_items
                        │
                        └── orders
                                │
                                └── table_sessions
                                        │
                                        └── tables

    table_sessions
            │
            └── payments

    daily_cash_drawer

------------------------------------------------------------------------

# Tecnologias Recomendadas

-   PostgreSQL
-   Backend: Go / Node / Java / Python
-   Frontend: Flutter / Web
-   Infraestrutura: Docker

------------------------------------------------------------------------

# Possíveis Melhorias Futuras

-   controle de usuários (garçom / admin)
-   integração com impressora da cozinha
-   relatórios de vendas
-   controle de estoque automático
-   dashboard financeiro
