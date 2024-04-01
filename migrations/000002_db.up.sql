CREATE TABLE IF NOT EXISTS posts (
    id UUID PRIMARY KEY NOT NULL,
    user_id UUID NOT NULL,
    content TEXT,
    title TEXT,
    likes BIGINT,
    dislikes BIGINT,
    views BIGINT,
    category TEXT,
    created_at TIMESTAMP,
    update_at TIMESTAMP,
    deleted_at TIMESTAMP
);