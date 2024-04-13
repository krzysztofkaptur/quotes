-- +goose Up
INSERT INTO authors (name) VALUES ('Julian'), ('Ricky');

-- +goose Down
DELETE FROM authors;
