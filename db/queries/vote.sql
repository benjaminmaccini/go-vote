-- name: CreateVote :one
INSERT INTO vote (id, candidate_id, rank, timestamp, voter_id, election_id)
VALUES (?, ?, ?, ?, ?, ?) RETURNING *;

-- name: UpdateVote :one
UPDATE vote
SET candidate_id = ?, rank = ?, timestamp = ?, voter_id = ?, election_id = ?
WHERE id = ?
RETURNING *;

-- name: DeleteVote :exec
DELETE FROM vote
WHERE id = ?;

-- name: ListVotes :many
SELECT * FROM vote
ORDER BY rank ASC, timestamp DESC
LIMIT ? OFFSET ?;
