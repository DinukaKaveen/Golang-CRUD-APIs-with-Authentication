-- name: GetAuthor :one
SELECT * FROM authors
WHERE id = ? 
LIMIT 1;

-- name: ListAuthors :many
SELECT * FROM authors
ORDER BY first_name;

-- name: CreateAuthor :execresult
INSERT INTO authors (first_name, last_name, bio)
VALUES (?, ?, ?);

-- name: UpdateAuthor :execresult
UPDATE authors
SET first_name = ?, last_name = ?, bio = ?
WHERE id = ?;

-- name: DeleteAuthor :exec
DELETE FROM authors
WHERE id = ?;