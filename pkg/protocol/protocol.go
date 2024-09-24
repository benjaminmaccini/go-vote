package protocol

import (
	"git.sr.ht/~bmaccini/go-vote/pkg/db"
)

type Candidate struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	ElectionID string `json:"election_id,omitempty"`
}

type Election struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ElectionResult struct {
	ID             string  `json:"id"`
	ElectionID     string  `json:"election_id,omitempty"`
	CandidateID    string  `json:"candidate_id,omitempty"`
	TotalVotes     int64   `json:"total_votes,omitempty"`
	VotePercentage float64 `json:"vote_percentage,omitempty"`
	Rank           int64   `json:"rank,omitempty"`
}

type Vote struct {
	ID          string `json:"id"`
	CandidateID string `json:"candidate_id,omitempty"`
	VoterID     string `json:"voter_id,omitempty"`
	Rank        int64  `json:"rank,omitempty"`
}

type Voter struct {
	ID  string `json:"id"`
	Zip string `json:"zip,omitempty"`
}

var ProtocolCommandMap = map[string]Protocol{
	"simpleMajority": new(SimpleMajority),
}

type Protocol interface {
	Init([]Candidate, *db.Queries) error
	Cast(Vote) error                      // Cast a vote
	GetID() string                        // Get the ID
	SetID(string)                         // Set the ID
	Results() (map[string]float64, error) // Compute the candidate's totals
	ValidateVote(Vote) (bool, error)      // Return if a cast vote is valid
}
