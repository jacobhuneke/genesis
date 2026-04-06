-- +goose Up
CREATE TABLE master(
    id UUID PRIMARY KEY,
    word TEXT NOT NULL UNIQUE,
    lemma TEXT NOT NULL,
    pos TEXT NOT NULL,
    etymology_id UUID NOT NULL,
    CONSTRAINT fk_etymologies_master
        FOREIGN KEY (etymology_id)
        REFERENCES etymologies(id)
);

-- +goose Down
DROP TABLE master;