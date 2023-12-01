
-- +migrate Up
CREATE TABLE transactions (
    id text PRIMARY KEY,
    user_id text,
    payment_id text,
    amount integer,
    transaction_type text,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp
);
-- +migrate Down

DROP TABLE "transactions";
