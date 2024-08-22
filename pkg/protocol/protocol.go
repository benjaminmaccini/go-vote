package protocol

import (
	"github.com/google/uuid"

	. "git.sr.ht/~bmaccini/go-vote/pkg/utils"
)

// This is our global election variable
var E Protocol

// Consumed by cmd/serve.go
var ProtocolCommandMap = map[string]Protocol{
	"simpleMajority": new(SimpleMajority),
}

type Candidate struct {
	Name string `json:"name"`
}

type Vote struct {
	Value     float64   `json:"value"`
	Candidate Candidate `json:"candidate"`
}

type Election struct {
	Id         string
	Name       string
	Candidates []Candidate
	Votes      []Vote
	Totals     map[string]float64 // Candidate name to their total
}

type Protocol interface {
	Init([]Candidate)
	Cast(Vote)                   // Cast a vote
	Tally()                      // Compute the candidate's totals
	Result() ([]string, float64) // Get the final result(s)
	Display()                    // Print the current totals
	GetId() string               // Get the id of the election
	SetId(string)                // Set the id, useful for testing
	ValidateVote(Vote) bool      // Return if a cast vote is valid
}

type SimpleMajority struct {
	Election
}

func (election *SimpleMajority) Init(candidates []Candidate) {
	InitLogger("INFO")
	election.Votes = make([]Vote, 0)
	election.Totals = make(map[string]float64)
	election.Candidates = candidates
	election.Name = "Simple Majority"
	id, _ := uuid.NewRandom()
	election.Id = id.String()
	Logger.Info("Election initialized.")
	election.Display()
}

func (election *SimpleMajority) Cast(vote Vote) {
	election.Votes = append(election.Votes, vote)
	Logger.Info("Ballot cast",
		"candidate", vote.Candidate,
		"value", vote.Value,
	)
}

func (election *SimpleMajority) Result() ([]string, float64) {
	election.Tally()
	var winners []string
	max := 0.0
	for candidateName, total := range election.Totals {
		if total > max {
			winners = []string{candidateName}
			max = total
		} else if total == max {
			winners = append(winners, candidateName)
		}
	}
	return winners, max
}

func (election *SimpleMajority) Tally() {
	election.Totals = make(map[string]float64)
	for _, vote := range election.Votes {
		election.Totals[vote.Candidate.Name] += vote.Value
	}
}

func (election *SimpleMajority) Display() {
	Logger.Info("Election display",
		"id", election.Id,
		"name", election.Name,
		"candidates", election.Candidates,
		"totals", election.Totals,
	)
}

func (election *SimpleMajority) GetId() string {
	return election.Id
}

func (election *SimpleMajority) SetId(id string) {
	election.Id = id
}

func (election *SimpleMajority) ValidateVote(vote Vote) bool {
	if vote.Value == 1 {
		for _, candidate := range election.Candidates {
			if vote.Candidate.Name == candidate.Name {
				return true
			}
		}
	}
	return false
}
