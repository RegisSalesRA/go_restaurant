CREATE TABLE IF NOT EXISTS daily_cash_drawer (
    id SERIAL PRIMARY KEY,
    opened_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    closed_at TIMESTAMP WITH TIME ZONE,
    initial_value INTEGER DEFAULT 0,   
    total_recorded INTEGER DEFAULT 0,  
    final_counted INTEGER DEFAULT 0,   
    status VARCHAR(10) DEFAULT 'aberto' CHECK (status IN ('aberto', 'fechado'))
);