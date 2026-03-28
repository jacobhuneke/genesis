
-- name: CreateLemma :one
INSERT INTO lemmas (id, word, lemma, pos, etymology_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
)
RETURNING *;

-- name: GetLemma :one
SELECT * FROM lemmas 
WHERE word = $1;