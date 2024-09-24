-- name: CreateElectionResult :one
INSERT INTO election_result (id, election_id, candidate_id, total_votes, vote_percentage, rank)
VALUES (?, ?, ?, ?, ?, ?) RETURNING *;

-- name: UpdateElectionResult :one
UPDATE election_result
SET election_id = ?, candidate_id = ?, total_votes = ?, vote_percentage = ?, rank = ?
WHERE id = ?
RETURNING *;

-- name: DeleteElectionResult :exec
DELETE FROM election_result
WHERE id = ?;

-- name: ListElectionResults :many
SELECT * FROM election_result
ORDER BY rank ASC
LIMIT ? OFFSET ?;
