
-- name: CreateEtymology :one
INSERT INTO etymologies (id, word, etymology, pos)
VALUES (
    $1,
    $2,
    $3,
    $4
)
RETURNING *;

-- name: GetEtymology :one
SELECT * FROM etymologies
WHERE word = $1;