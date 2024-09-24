-- name: CreateVoter :one
INSERT INTO voter (id, zip)
VALUES (?, ?) RETURNING *;

-- name: UpdateVoter :one
UPDATE voter
SET zip = ?
WHERE id = ?
RETURNING *;

-- name: DeleteVoter :exec
DELETE FROM voter
WHERE id = ?;

-- name: ListVoters :many
SELECT * FROM voter
ORDER BY id
LIMIT ? OFFSET ?;
