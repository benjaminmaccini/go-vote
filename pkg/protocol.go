package protocol

import (
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

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
	Tally      func()
}

type Protocol interface {
	Init([]Candidate)
	Cast(Vote)                      // Cast a vote
	tally()                         // Closure for compute the candidate's totals, without recomputing
	Result() ([]Candidate, float64) // Get the final result(s)
	Display()                       // Print the current totals
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
	election.Tally = election.tally()
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

func (election *SimpleMajority) tally() func() {
	previous := 0
	return func() {
		for i, vote := range election.Votes[previous+1:] {
			if _, exists := election.Totals[vote.Candidate.Name]; exists {
				election.Totals[vote.Candidate.Name] += vote.Value
			} else {
				election.Totals[vote.Candidate.Name] = vote.Value
			}
			previous = i
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
