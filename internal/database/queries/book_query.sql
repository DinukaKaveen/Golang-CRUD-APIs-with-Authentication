-- name: GetBook :one
SELECT * FROM books
WHERE id = ? 
LIMIT 1;

-- name: ListBooks :many
SELECT * FROM books
ORDER BY title;

-- name: CreateBook :execresult
INSERT INTO books (title, author_id)
VALUES (?, ?);

-- name: UpdateBook :execresult
UPDATE books
SET title = ?, author_id = ?
WHERE id = ?;

-- name: DeleteBook :exec
DELETE FROM books
WHERE id = ?;