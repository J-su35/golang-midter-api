-- +goose Up
ALTER TABLE items ADD status text NOT NULL;
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
ALTER TABLE items DROP COLUMN status;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
