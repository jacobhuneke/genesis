

-- name: CreateMaster :one
INSERT INTO master (id, word, lemma, pos, etymology_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
)
RETURNING *;

-- name: GetMaster :one
SELECT * FROM master
WHERE word = $1;