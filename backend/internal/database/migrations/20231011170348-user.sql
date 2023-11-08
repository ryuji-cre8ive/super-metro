
-- +migrate Up
CREATE TABLE "users" (
    id text PRIMARY KEY,
    name text,
    email text,
    password text,
    session_token text,
    valance integer,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp
);
-- +migrate Down
DROP TABLE "users";