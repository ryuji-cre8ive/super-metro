
-- +migrate Up
CREATE TABLE "users" (
    id text PRIMARY KEY,
    name text,
    email text,
    password text,
    session_token text
);
-- +migrate Down
DROP TABLE users;