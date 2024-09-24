package protocol

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"git.sr.ht/~bmaccini/go-vote/pkg/db"
	. "git.sr.ht/~bmaccini/go-vote/pkg/utils"
)

type SimpleMajority struct {
	Election
	*db.Queries
}

func (sm *SimpleMajority) Init(candidates []Candidate, queries *db.Queries) error {
	sm.Election = Election{
		ID:   GetStringUUIDv7(),
		Name: "Simple Majority",
	}
	sm.Queries = queries

	_, err := sm.Queries.CreateElection(context.Background(), db.CreateElectionParams{
		ID:   sm.Election.ID,
		Name: sm.Election.Name,
	})
	if err != nil {
		return err
	}

	for _, c := range candidates {
		_, err = sm.Queries.CreateCandidate(context.Background(), db.CreateCandidateParams{
			ID:         c.ID,
			Name:       c.Name,
			ElectionID: sql.NullString{String: sm.Election.ID, Valid: true},
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func (sm *SimpleMajority) Cast(vote Vote) error {
	_, err := sm.Queries.CreateVote(context.Background(), db.CreateVoteParams{
		ID:          vote.ID,
		CandidateID: vote.CandidateID,
		Rank:        sql.NullInt64{Int64: 1, Valid: true},
		Timestamp:   time.Now(),
		VoterID:     vote.VoterID,
	})
	if err != nil {
		return err
	}
	return nil
}

func (sm *SimpleMajority) Results() (map[string]float64, error) {
	votes := []db.Vote{}
	var offset int64 = 0
	for {
		vs, err := sm.Queries.ListVotes(context.Background(), db.ListVotesParams{
			Limit:  PAGING_LIMIT,
			Offset: offset,
		})
		if err != nil {
			return nil, fmt.Errorf("error listing votes: %w", err)
		}
		if len(vs) == 0 {
			break
		}
		votes = append(votes, vs...)
		offset += PAGING_LIMIT
	}

	totals := make(map[string]float64)
	for _, vote := range votes {
		if _, exists := totals[vote.CandidateID]; !exists {
			totals[vote.CandidateID] = 0
		}
		totals[vote.CandidateID] += 1
	}
	return totals, nil
}

func (sm *SimpleMajority) GetID() string {
	return sm.Election.ID
}

func (sm *SimpleMajority) SetID(id string) {
	sm.Election.ID = id
}

func (sm *SimpleMajority) ValidateVote(vote Vote) (bool, error) {
	_, err := sm.Queries.GetCandidate(context.Background(), vote.CandidateID)
	if err != nil {
		return true, nil
	}
	return false, nil
}
