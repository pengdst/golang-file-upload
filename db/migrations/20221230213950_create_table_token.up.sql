CREATE TABLE IF NOT EXISTS access_tokens
(
    id         bigserial primary key,
    user_id    bigint constraint fk_access_tokens_user references users,
    auth_token text,
    role       text,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);

CREATE INDEX idx_access_tokens_deleted_at on access_tokens (deleted_at);

