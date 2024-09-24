-- name: CreateCandidate :one
INSERT INTO candidate (id, name, election_id)
VALUES (?, ?, ?) RETURNING *;

-- name: UpdateCandidate :one
UPDATE candidate
SET name = ?, election_id = ?
WHERE id = ?
RETURNING *;

-- name: DeleteCandidate :exec
DELETE FROM candidate
WHERE id = ?;

-- name: ListCandidates :many
SELECT * FROM candidate
ORDER BY name
LIMIT ? OFFSET ?;

-- name: GetCandidate :one
SELECT * FROM candidate
WHERE id = ?;

-- name: GetCandidatesByElection :many
SELECT * FROM candidate
WHERE election_id = ?
ORDER BY name;
