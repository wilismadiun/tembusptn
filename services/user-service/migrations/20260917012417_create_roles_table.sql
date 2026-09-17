-- +goose Up
CREATE TABLE roles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE
);

INSERT INTO roles (name, id) VALUES
    ('admin', 1),
    ('student', 2),
    ('teacher', 3);

-- +goose Down
DROP TABLE roles;