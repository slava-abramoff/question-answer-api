-- +goose Up
ALTER TABLE questions ADD COLUMN updated_at TIMESTAMP NOT NULL DEFAULT now();
ALTER TABLE questions ADD COLUMN deleted_at TIMESTAMP;

ALTER TABLE answers ADD COLUMN updated_at TIMESTAMP NOT NULL DEFAULT now();
ALTER TABLE answers ADD COLUMN deleted_at TIMESTAMP;

-- +goose Down
ALTER TABLE questions DROP COLUMN updated_at;
ALTER TABLE questions DROP COLUMN deleted_at;

ALTER TABLE answers DROP COLUMN updated_at;
ALTER TABLE answers DROP COLUMN deleted_at;
