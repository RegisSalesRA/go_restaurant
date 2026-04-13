CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    category_id INTEGER REFERENCES categories(id) ON DELETE SET NULL,
    name VARCHAR(100) NOT NULL,
    price INTEGER NOT NULL, 
    stock_quantity INTEGER DEFAULT 0,
    active BOOLEAN DEFAULT TRUE
);