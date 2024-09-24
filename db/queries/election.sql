-- name: CreateElection :one
INSERT INTO election (id, name)
VALUES (?, ?) RETURNING *;

-- name: UpdateElection :one
UPDATE election
SET name = ?
WHERE id = ?
RETURNING *;

-- name: DeleteElection :exec
DELETE FROM election
WHERE id = ?;

-- name: ListElections :many
SELECT * FROM election
ORDER BY name
LIMIT ? OFFSET ?;
