-- +goose Up
CREATE TABLE subscriptions (
    id VARCHAR PRIMARY KEY,
    name VARCHAR NOT NULL,
    price BIGINT NOT NULL
);

-- +goose Down
DROP TABLE subscriptions;