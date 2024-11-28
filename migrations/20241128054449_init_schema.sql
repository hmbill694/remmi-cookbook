-- +goose Up
-- +goose StatementBegin
-- 1. Create users table
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username TEXT NOT NULL UNIQUE,
    email TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW ()
);

-- 2. Create recipes table
CREATE TABLE IF NOT EXISTS recipes (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    title TEXT NOT NULL,
    description TEXT,
    created_at TIMESTAMPTZ DEFAULT NOW (),
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

-- 3. Create recipe_chunks table
CREATE TABLE IF NOT EXISTS recipe_chunks (
    id SERIAL PRIMARY KEY,
    recipe_id INTEGER NOT NULL,
    chunk_order INTEGER NOT NULL,
    content TEXT NOT NULL,
    FOREIGN KEY (recipe_id) REFERENCES recipes (id) ON DELETE CASCADE,
    UNIQUE (recipe_id, chunk_order)
);

-- 4. Create chats table
CREATE TABLE IF NOT EXISTS chats (
    id SERIAL PRIMARY KEY,
    recipe_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    started_at TIMESTAMPTZ DEFAULT NOW (),
    FOREIGN KEY (recipe_id) REFERENCES recipes (id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

-- 5. Create chat_messages table
CREATE TABLE IF NOT EXISTS chat_messages (
    id SERIAL PRIMARY KEY,
    chat_id INTEGER NOT NULL,
    sender TEXT NOT NULL, -- 'user' or 'llm'
    message TEXT NOT NULL,
    sent_at TIMESTAMPTZ DEFAULT NOW (),
    FOREIGN KEY (chat_id) REFERENCES chats (id) ON DELETE CASCADE
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS chat_messages;

DROP TABLE IF EXISTS chats;

DROP TABLE IF EXISTS recipe_chunks;

DROP TABLE IF EXISTS recipes;

-- +goose StatementEnd
