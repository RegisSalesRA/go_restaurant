CREATE TABLE payments (
    id SERIAL PRIMARY KEY,
    session_id UUID REFERENCES table_sessions(id) ON DELETE CASCADE,
    amount INTEGER NOT NULL,
    payment_method VARCHAR(20) CHECK (payment_method IN ('dinheiro', 'cartao', 'pix')),
    paid_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);