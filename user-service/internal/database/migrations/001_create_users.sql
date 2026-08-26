-- +goose Up

CREATE TABLE users
(
    id       BIGSERIAL PRIMARY KEY,
    username TEXT NOT NULL,
    email    TEXT NOT NULL UNIQUE
);

-- +goose Down

DROP TABLE users;