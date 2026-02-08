-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users(
    id BIGINT NOT NULL PRIMARY KEY,
    tg_user_id BIGINT UNIQUE NOT NULL,
    name VARCHAR(64) NOT NULL,
    role VARCHAR(16) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NULL
);

COMMENT ON COLUMN users.id IS 'ID пользователя';
COMMENT ON COLUMN users.tg_user_id IS 'ID tg';
COMMENT ON COLUMN users.name IS 'Имя пользователя';
COMMENT ON COLUMN users.role IS 'Роль ';
COMMENT ON COLUMN users.created_at IS 'Дата регистрации';
COMMENT ON COLUMN users.updated_at IS 'Дата изменения';
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users CASCADE;
-- +goose StatementEnd