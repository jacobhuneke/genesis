
-- +goose Up
CREATE TABLE etymologies(
    id UUID PRIMARY KEY,
    word TEXT NOT NULL UNIQUE,
    etymology TEXT NOT NULL,
    pos TEXT NOT NULL
);

-- +goose Down
DROP TABLE etymologies;