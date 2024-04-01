CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY NOT NULL,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    bio TEXT,
    website VARCHAR(255),
    deleted_at TIMESTAMP
);
