-- ==========================================
-- 1. CATEGORIAS E PRODUTOS (MENU)
-- ==========================================
CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);

CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    category_id INTEGER REFERENCES categories(id),
    name VARCHAR(100) NOT NULL,
    price INTEGER NOT NULL, -- Preço em centavos (evita problemas de precisão decimal)
    stock_quantity INTEGER DEFAULT 0
);

-- ==========================================
-- 2. INFRAESTRUTURA (MESAS)
-- ==========================================
CREATE TABLE tables (
    id SERIAL PRIMARY KEY,
    number INTEGER UNIQUE NOT NULL,
    status VARCHAR(20) DEFAULT 'disponivel' -- 'disponivel', 'ocupada', 'suja'
);

-- ==========================================
-- 3. OPERAÇÃO (SESSÕES E PEDIDOS)
-- ==========================================
-- Uma sessão é aberta quando o cliente senta à mesa
CREATE TABLE table_sessions (
    id UUID PRIMARY KEY,
    table_id INTEGER REFERENCES tables(id),
    opened_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    closed_at TIMESTAMP WITH TIME ZONE, -- Nulo enquanto a conta estiver aberta
    total_amount INTEGER DEFAULT 0      -- Soma total de todos os itens consumidos
);

-- Pedidos individuais feitos durante a permanência na mesa
CREATE TABLE orders (
    id UUID PRIMARY KEY,
    session_id UUID REFERENCES table_sessions(id),
    status VARCHAR(20) DEFAULT 'pendente', -- 'preparando', 'pronto', 'entregue'
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Itens vinculados a cada pedido específico
CREATE TABLE order_items (
    id SERIAL PRIMARY KEY,
    order_id UUID REFERENCES orders(id) ON DELETE CASCADE,
    product_id INTEGER REFERENCES products(id),
    quantity INTEGER NOT NULL,
    unit_price INTEGER NOT NULL -- Preço fixado no momento do pedido (histórico)
);

-- ==========================================
-- 4. FINANCEIRO (PAGAMENTOS E CAIXA)
-- ==========================================
-- Registro de como a conta foi paga
CREATE TABLE payments (
    id SERIAL PRIMARY KEY,
    session_id UUID REFERENCES table_sessions(id),
    amount INTEGER NOT NULL,            -- Valor pago em centavos
    payment_method VARCHAR(20),         -- 'dinheiro', 'cartao', 'pix'
    paid_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Fluxo de caixa diário do restaurante
CREATE TABLE daily_cash_drawer (
    id SERIAL PRIMARY KEY,
    opened_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    closed_at TIMESTAMP WITH TIME ZONE,
    initial_value INTEGER DEFAULT 0,    -- Dinheiro inicial para troco
    final_value INTEGER DEFAULT 0,      -- Valor total acumulado no dia
    status VARCHAR(10) DEFAULT 'aberto' -- 'aberto', 'fechado'
);

-- ==========================================
-- DADOS PARA TESTE INICIAL (OPCIONAL)
-- ==========================================
INSERT INTO categories (name) VALUES ('Bebidas'), ('Pratos Principais');
INSERT INTO products (name, category_id, price, stock_quantity) VALUES 
('Coca-Cola 350ml', 1, 600, 50),
('Hambúrguer Gourmet', 2, 3500, 20);
INSERT INTO tables (number, status) VALUES (1, 'disponivel'), (2, 'disponivel'), (3, 'disponivel');
