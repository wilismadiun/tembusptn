-- +goose Up
CREATE TABLE users (
    id VARCHAR PRIMARY KEY,
    email VARCHAR NOT NULL UNIQUE,
    name VARCHAR NOT NULL,
    phone VARCHAR,
    password VARCHAR NOT NULL
);

-- +goose Down
DROP TABLE users;