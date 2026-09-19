-- +goose Up
CREATE TABLE user_subscriptions (
    id VARCHAR(255) PRIMARY KEY,
    user_id VARCHAR(255) NOT NULL,
    subscription_id VARCHAR(255) NOT NULL,
    start_at timestamptz NOT NULL,
    end_at timestamptz NOT NULL,
    status VARCHAR(50) NOT NULL,
    CONSTRAINT fk_subscription FOREIGN KEY (subscription_id) REFERENCES subscriptions(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE IF EXISTS user_subscriptions;