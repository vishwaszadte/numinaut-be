-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS expressions (
    id SERIAL PRIMARY KEY,
    uuid UUID NOT NULL UNIQUE,
    expression VARCHAR(255) NOT NULL UNIQUE,
    result REAL NOT NULL,
    num_operands INT NOT NULL,
    difficulty INT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS expressions;
-- +goose StatementEnd
