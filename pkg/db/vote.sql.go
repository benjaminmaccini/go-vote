// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: vote.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createVote = `-- name: CreateVote :one
INSERT INTO vote (id, candidate_id, rank, timestamp, voter_id, election_id)
VALUES (?, ?, ?, ?, ?, ?) RETURNING id, election_id, candidate_id, rank, timestamp, voter_id
`

type CreateVoteParams struct {
	ID          string        `json:"id"`
	CandidateID string        `json:"candidate_id"`
	Rank        sql.NullInt64 `json:"rank"`
	Timestamp   time.Time     `json:"timestamp"`
	VoterID     string        `json:"voter_id"`
	ElectionID  string        `json:"election_id"`
}

func (q *Queries) CreateVote(ctx context.Context, arg CreateVoteParams) (Vote, error) {
	row := q.db.QueryRowContext(ctx, createVote,
		arg.ID,
		arg.CandidateID,
		arg.Rank,
		arg.Timestamp,
		arg.VoterID,
		arg.ElectionID,
	)
	var i Vote
	err := row.Scan(
		&i.ID,
		&i.ElectionID,
		&i.CandidateID,
		&i.Rank,
		&i.Timestamp,
		&i.VoterID,
	)
	return i, err
}

const deleteVote = `-- name: DeleteVote :exec
DELETE FROM vote
WHERE id = ?
`

func (q *Queries) DeleteVote(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteVote, id)
	return err
}

const listVotes = `-- name: ListVotes :many
SELECT id, election_id, candidate_id, rank, timestamp, voter_id FROM vote
ORDER BY rank ASC, timestamp DESC
LIMIT ? OFFSET ?
`

type ListVotesParams struct {
	Limit  int64 `json:"limit"`
	Offset int64 `json:"offset"`
}

func (q *Queries) ListVotes(ctx context.Context, arg ListVotesParams) ([]Vote, error) {
	rows, err := q.db.QueryContext(ctx, listVotes, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Vote
	for rows.Next() {
		var i Vote
		if err := rows.Scan(
			&i.ID,
			&i.ElectionID,
			&i.CandidateID,
			&i.Rank,
			&i.Timestamp,
			&i.VoterID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateVote = `-- name: UpdateVote :one
UPDATE vote
SET candidate_id = ?, rank = ?, timestamp = ?, voter_id = ?, election_id = ?
WHERE id = ?
RETURNING id, election_id, candidate_id, rank, timestamp, voter_id
`

type UpdateVoteParams struct {
	CandidateID string        `json:"candidate_id"`
	Rank        sql.NullInt64 `json:"rank"`
	Timestamp   time.Time     `json:"timestamp"`
	VoterID     string        `json:"voter_id"`
	ElectionID  string        `json:"election_id"`
	ID          string        `json:"id"`
}

func (q *Queries) UpdateVote(ctx context.Context, arg UpdateVoteParams) (Vote, error) {
	row := q.db.QueryRowContext(ctx, updateVote,
		arg.CandidateID,
		arg.Rank,
		arg.Timestamp,
		arg.VoterID,
		arg.ElectionID,
		arg.ID,
	)
	var i Vote
	err := row.Scan(
		&i.ID,
		&i.ElectionID,
		&i.CandidateID,
		&i.Rank,
		&i.Timestamp,
		&i.VoterID,
	)
	return i, err
}
