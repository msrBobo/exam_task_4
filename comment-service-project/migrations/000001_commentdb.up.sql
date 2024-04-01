CREATE TABLE IF NOT EXISTS comments (
    id UUID PRIMARY KEY NOT NULL,
    post_id UUID NOT NULL,
    user_id UUID NOT NULL,
    content TEXT NOT NULL,
    likes BIGINT DEFAULT 0,
    dislikes BIGINT DEFAULT 0,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);
