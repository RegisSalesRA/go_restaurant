
CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE tables (
    id SERIAL PRIMARY KEY,
    number INTEGER UNIQUE NOT NULL,
    status VARCHAR(20) DEFAULT 'disponivel' CHECK (status IN ('disponivel', 'ocupada', 'suja'))
);


CREATE TABLE daily_cash_drawer (
    id SERIAL PRIMARY KEY,
    opened_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    closed_at TIMESTAMP WITH TIME ZONE,
    initial_value INTEGER DEFAULT 0,   
    total_recorded INTEGER DEFAULT 0,  
    final_counted INTEGER DEFAULT 0,   
    status VARCHAR(10) DEFAULT 'aberto' CHECK (status IN ('aberto', 'fechado'))
);


CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    category_id INTEGER REFERENCES categories(id) ON DELETE SET NULL,
    name VARCHAR(100) NOT NULL,
    price INTEGER NOT NULL, 
    stock_quantity INTEGER DEFAULT 0,
    active BOOLEAN DEFAULT TRUE
);


CREATE TABLE table_sessions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    table_id INTEGER REFERENCES tables(id) ON DELETE CASCADE,
    opened_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    closed_at TIMESTAMP WITH TIME ZONE,
    total_amount INTEGER DEFAULT 0
);


CREATE TABLE orders (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    session_id UUID REFERENCES table_sessions(id) ON DELETE CASCADE,
    status VARCHAR(20) DEFAULT 'pendente' CHECK (status IN ('pendente', 'preparando', 'pronto', 'entregue')),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE order_items (
    id SERIAL PRIMARY KEY,
    order_id UUID REFERENCES orders(id) ON DELETE CASCADE,
    product_id INTEGER REFERENCES products(id),
    quantity INTEGER NOT NULL CHECK (quantity > 0),
    unit_price INTEGER NOT NULL 
);


CREATE TABLE payments (
    id SERIAL PRIMARY KEY,
    session_id UUID REFERENCES table_sessions(id) ON DELETE CASCADE,
    amount INTEGER NOT NULL,
    payment_method VARCHAR(20) CHECK (payment_method IN ('dinheiro', 'cartao', 'pix')),
    paid_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);


