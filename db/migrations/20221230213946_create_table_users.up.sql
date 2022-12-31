CREATE TABLE IF NOT EXISTS users
(
    id         bigserial primary key,
    name       text,
    email      text,
    role      text,
    password   text,
    deleted_at timestamp with time zone,
    updated_at timestamp with time zone,
    created_at timestamp with time zone
);

CREATE INDEX idx_users_deleted_at ON users (deleted_at);
CREATE UNIQUE INDEX idx_users_email ON users (email);