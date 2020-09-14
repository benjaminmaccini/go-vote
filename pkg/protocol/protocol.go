package protocol

import (
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

// Consumed by cmd/serve.go
var ProtocolCommandMap = map[string]Protocol{
	"simpleMajority": new(SimpleMajority),
}

type Candidate struct {
	Name string
}

type Vote struct {
	Value     float64
	Candidate Candidate
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
}

type SimpleMajority struct {
	Election
}

func (election *SimpleMajority) Init(candidates []Candidate) {
	election.Votes = make([]Vote, 0)
	election.Totals = make(map[string]float64)
	election.Candidates = candidates
	election.Name = "Simple Majority"
	id, _ := uuid.NewRandom()
	election.Id = id.String()
	log.Info("Election initialized.")
	election.Display()
}

func (election *SimpleMajority) Cast(vote Vote) {
	election.Votes = append(election.Votes, vote)
	log.WithFields(log.Fields{
		"candidate": vote.Candidate,
		"value":     vote.Value,
	}).Info("Vote cast")
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
		if _, exists := election.Totals[vote.Candidate.Name]; exists {
			election.Totals[vote.Candidate.Name] += vote.Value
		} else {
			election.Totals[vote.Candidate.Name] = vote.Value
		}
	}
}

func (election *SimpleMajority) Display() {
	log.WithFields(log.Fields{
		"id":         election.Id,
		"name":       election.Name,
		"candidates": election.Candidates,
		"totals":     election.Totals,
	}).Info()
}

func (election *SimpleMajority) GetId() string {
	return election.Id
}
