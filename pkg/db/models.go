// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"database/sql"
	"time"
)

type Candidate struct {
	ID         string         `json:"id"`
	Name       string         `json:"name"`
	ElectionID sql.NullString `json:"election_id"`
}

type Election struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ElectionResult struct {
	ID             string          `json:"id"`
	ElectionID     sql.NullString  `json:"election_id"`
	CandidateID    sql.NullString  `json:"candidate_id"`
	TotalVotes     sql.NullInt64   `json:"total_votes"`
	VotePercentage sql.NullFloat64 `json:"vote_percentage"`
	Rank           sql.NullInt64   `json:"rank"`
}

type SchemaMigrations struct {
	Version string `json:"version"`
}

type Vote struct {
	ID          string        `json:"id"`
	ElectionID  string        `json:"election_id"`
	CandidateID string        `json:"candidate_id"`
	Rank        sql.NullInt64 `json:"rank"`
	Timestamp   time.Time     `json:"timestamp"`
	VoterID     string        `json:"voter_id"`
}

type Voter struct {
	ID  string         `json:"id"`
	Zip sql.NullString `json:"zip"`
}
