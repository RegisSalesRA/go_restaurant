-- ============================================================
-- 1. MENU: CATEGORIAS E PRODUTOS
-- ============================================================
-- Armazena as categorias do cardápio (ex: Bebidas, Sobremesas)
CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);

-- Armazena os itens disponíveis para venda
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    category_id INTEGER REFERENCES categories(id),
    name VARCHAR(100) NOT NULL,
    price INTEGER NOT NULL, -- Valor em centavos para evitar erros de ponto flutuante
    stock_quantity INTEGER DEFAULT 0
);

-- ============================================================
-- 2. INFRAESTRUTURA: MESAS
-- ============================================================
-- Controle físico das mesas do estabelecimento
CREATE TABLE tables (
    id SERIAL PRIMARY KEY,
    number INTEGER UNIQUE NOT NULL,
    status VARCHAR(20) DEFAULT 'disponivel' -- status: 'disponivel', 'ocupada', 'suja'
);

-- ============================================================
-- 3. OPERAÇÃO: SESSÕES E PEDIDOS
-- ============================================================
-- Representa a estadia de um cliente em uma mesa (a "conta" da mesa)
CREATE TABLE table_sessions (
    id UUID PRIMARY KEY,
    table_id INTEGER REFERENCES tables(id),
    opened_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    closed_at TIMESTAMP WITH TIME ZONE, -- Se preenchido, a mesa foi liberada
    total_amount INTEGER DEFAULT 0      -- Soma acumulada da sessão
);

-- Registra cada vez que o garçom faz um pedido para a cozinha
CREATE TABLE orders (
    id UUID PRIMARY KEY,
    session_id UUID REFERENCES table_sessions(id),
    status VARCHAR(20) DEFAULT 'pendente', -- 'preparando', 'pronto', 'entregue'
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Detalhamento de cada item dentro de um pedido
CREATE TABLE order_items (
    id SERIAL PRIMARY KEY,
    order_id UUID REFERENCES orders(id) ON DELETE CASCADE,
    product_id INTEGER REFERENCES products(id),
    quantity INTEGER NOT NULL,
    unit_price INTEGER NOT NULL -- Preço capturado no momento da venda (histórico)
);

-- ============================================================
-- 4. FINANCEIRO: PAGAMENTOS E CAIXA
-- ============================================================
-- Registro das transações financeiras realizadas nas sessões
CREATE TABLE payments (
    id SERIAL PRIMARY KEY,
    session_id UUID REFERENCES table_sessions(id),
    amount INTEGER NOT NULL,            -- Valor pago em centavos
    payment_method VARCHAR(20),         -- 'dinheiro', 'cartao', 'pix'
    paid_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Controle do fluxo de caixa diário do restaurante
CREATE TABLE daily_cash_drawer (
    id SERIAL PRIMARY KEY,
    opened_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    closed_at TIMESTAMP WITH TIME ZONE,
    initial_value INTEGER DEFAULT 0,    -- Fundo de caixa (troco)
    final_value INTEGER DEFAULT 0,      -- Total final após fechamento
    status VARCHAR(10) DEFAULT 'aberto' -- 'aberto' ou 'fechado'
);

-- ============================================================
-- DADOS PARA TESTE INICIAL
-- ============================================================
INSERT INTO categories (name) VALUES ('Bebidas'), ('Pratos Principais');

INSERT INTO products (name, category_id, price, stock_quantity) VALUES 
('Coca-Cola 350ml', 1, 600, 50),
('Hambúrguer Gourmet', 2, 3500, 20);

INSERT INTO tables (number, status) VALUES (1, 'disponivel'), (2, 'disponivel'), (3, 'disponivel');
