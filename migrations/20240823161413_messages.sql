-- +goose Up
CREATE TABLE messages
(
    id           SERIAL PRIMARY KEY,
    message_from TEXT,
    message_text TEXT,
    created_at   TIMESTAMP NOT NULL DEFAULT now(),
    updated_at   TIMESTAMP
);

-- +goose Down
DROP TABLE IF EXISTS messages;
