// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: candidate.sql

package db

import (
	"context"
	"database/sql"
)

const createCandidate = `-- name: CreateCandidate :one
INSERT INTO candidate (id, name, election_id)
VALUES (?, ?, ?) RETURNING id, name, election_id
`

type CreateCandidateParams struct {
	ID         string         `json:"id"`
	Name       string         `json:"name"`
	ElectionID sql.NullString `json:"election_id"`
}

func (q *Queries) CreateCandidate(ctx context.Context, arg CreateCandidateParams) (Candidate, error) {
	row := q.db.QueryRowContext(ctx, createCandidate, arg.ID, arg.Name, arg.ElectionID)
	var i Candidate
	err := row.Scan(&i.ID, &i.Name, &i.ElectionID)
	return i, err
}

const deleteCandidate = `-- name: DeleteCandidate :exec
DELETE FROM candidate
WHERE id = ?
`

func (q *Queries) DeleteCandidate(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteCandidate, id)
	return err
}

const getCandidate = `-- name: GetCandidate :one
SELECT id, name, election_id FROM candidate
WHERE id = ?
`

func (q *Queries) GetCandidate(ctx context.Context, id string) (Candidate, error) {
	row := q.db.QueryRowContext(ctx, getCandidate, id)
	var i Candidate
	err := row.Scan(&i.ID, &i.Name, &i.ElectionID)
	return i, err
}

const getCandidatesByElection = `-- name: GetCandidatesByElection :many
SELECT id, name, election_id FROM candidate
WHERE election_id = ?
ORDER BY name
`

func (q *Queries) GetCandidatesByElection(ctx context.Context, electionID sql.NullString) ([]Candidate, error) {
	rows, err := q.db.QueryContext(ctx, getCandidatesByElection, electionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Candidate
	for rows.Next() {
		var i Candidate
		if err := rows.Scan(&i.ID, &i.Name, &i.ElectionID); err != nil {
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

const listCandidates = `-- name: ListCandidates :many
SELECT id, name, election_id FROM candidate
ORDER BY name
LIMIT ? OFFSET ?
`

type ListCandidatesParams struct {
	Limit  int64 `json:"limit"`
	Offset int64 `json:"offset"`
}

func (q *Queries) ListCandidates(ctx context.Context, arg ListCandidatesParams) ([]Candidate, error) {
	rows, err := q.db.QueryContext(ctx, listCandidates, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Candidate
	for rows.Next() {
		var i Candidate
		if err := rows.Scan(&i.ID, &i.Name, &i.ElectionID); err != nil {
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

const updateCandidate = `-- name: UpdateCandidate :one
UPDATE candidate
SET name = ?, election_id = ?
WHERE id = ?
RETURNING id, name, election_id
`

type UpdateCandidateParams struct {
	Name       string         `json:"name"`
	ElectionID sql.NullString `json:"election_id"`
	ID         string         `json:"id"`
}

func (q *Queries) UpdateCandidate(ctx context.Context, arg UpdateCandidateParams) (Candidate, error) {
	row := q.db.QueryRowContext(ctx, updateCandidate, arg.Name, arg.ElectionID, arg.ID)
	var i Candidate
	err := row.Scan(&i.ID, &i.Name, &i.ElectionID)
	return i, err
}
