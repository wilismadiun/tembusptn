-- +goose Up
CREATE TABLE subscriptions (
    id VARCHAR PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR NOT NULL,
    price BIGINT NOT NULL
);

-- +goose Down
DROP TABLE subscriptions;