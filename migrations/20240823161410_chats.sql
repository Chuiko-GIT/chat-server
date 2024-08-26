-- +goose Up
CREATE TABLE chats
(
    id         SERIAL PRIMARY KEY,
    usernames  TEXT[],
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP
);

-- +goose Down
DROP TABLE IF EXISTS chats;
