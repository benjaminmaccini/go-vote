-- name: CreateVoter :one
INSERT INTO voter (id, zip)
VALUES (?, ?) RETURNING *;

-- name: UpdateVoter :one
UPDATE voter
SET zip = ?
WHERE id = ?
RETURNING *;

-- name: UpsertVoter :one
INSERT INTO voter (id, zip)
VALUES (?, ?)
ON CONFLICT (id) DO UPDATE SET
    zip = excluded.zip
RETURNING *;

-- name: DeleteVoter :exec
DELETE FROM voter
WHERE id = ?;

-- name: ListVoters :many
SELECT * FROM voter
ORDER BY id
LIMIT ? OFFSET ?;

-- name: GetVoterByIdIfVoted :one
SELECT voter.*
FROM voter
INNER JOIN vote ON voter.id = vote.voter_id
WHERE voter.id = ? AND vote.election_id = ?
LIMIT 1;
