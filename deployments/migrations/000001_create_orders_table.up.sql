-- Using BIGSERIAL for auto-incrementing uint64 IDs
CREATE TABLE IF NOT EXISTS orders (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    status SMALLINT NOT NULL DEFAULT 0,
    total_price BIGINT NOT NULL DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE -- Include if your BaseModel uses GORM soft deletes
);

-- Indexing UserID as specified in your GORM tag
CREATE INDEX IF NOT EXISTS idx_orders_user_id ON orders(user_id);