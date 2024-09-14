-- +goose Up
ALTER TABLE items ADD amount real NOT NULL;
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
ALTER TABLE items DROP COLUMN amount;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd