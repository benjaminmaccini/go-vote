package protocol

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"git.sr.ht/~bmaccini/go-vote/pkg/db"
	"git.sr.ht/~bmaccini/go-vote/pkg/utils"
	. "git.sr.ht/~bmaccini/go-vote/pkg/utils"
)

type RankedChoice struct {
	Election
	*db.Queries
}

func (sm *RankedChoice) Init(candidates []Candidate, queries *db.Queries) error {
	sm.Election = Election{
		ID:   GetStringUUIDv7(),
		Name: "RankedChoice",
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

func (sm *RankedChoice) Cast(votes []Vote) error {
	for _, vote := range votes {
		_, err := sm.Queries.CreateVote(context.Background(), db.CreateVoteParams{
			ID:          vote.ID,
			CandidateID: vote.CandidateID,
			Rank:        sql.NullInt64{Int64: vote.Rank, Valid: true},
			Timestamp:   time.Now(),
			VoterID:     vote.VoterID,
			ElectionID:  vote.ElectionID,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (sm *RankedChoice) Results() (map[string]float64, error) {
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
		totals[vote.CandidateID] += float64(vote.Rank.Int64)
	}
	return totals, nil
}

func (sm *RankedChoice) GetID() string {
	return sm.Election.ID
}

func (sm *RankedChoice) SetID(id string) {
	sm.Election.ID = id
}

func (sm *RankedChoice) ValidateVote(votes []Vote) (bool, error) {
	// Create a map to track ranks, candidates, voterIDs
	rankMap := make(map[int64]bool)
	candidateMap := make(map[string]bool)
	voterIDs := []string{}

	// Check the number of votes matches the number of candidates
	candidates := []db.Candidate{}
	var offset int64 = 0
	for {
		cs, err := sm.Queries.ListCandidates(context.Background(), db.ListCandidatesParams{
			Limit:  PAGING_LIMIT,
			Offset: offset,
		})
		if err != nil {
			return false, fmt.Errorf("error listing candidates: %w", err)
		}
		if len(cs) == 0 {
			break
		}
		candidates = append(candidates, cs...)
		offset += PAGING_LIMIT
	}
	if len(votes) != len(candidates) {
		return false, fmt.Errorf("RankedChoice: number of votes does not match number of candidates")
	}

	for _, vote := range votes {
		// Check if the candidate has already been voted for
		if candidateMap[vote.CandidateID] {
			return false, fmt.Errorf("candidate %s received multiple votes", vote.CandidateID)
		}
		candidateMap[vote.CandidateID] = true

		// Check if the rank has already been used
		if rankMap[vote.Rank] {
			return false, fmt.Errorf("rank %d used multiple times", vote.Rank)
		}
		rankMap[vote.Rank] = true

		// Check if the rank is within the valid range
		if vote.Rank < 1 || vote.Rank > int64(len(votes)) {
			return false, fmt.Errorf("invalid rank %d", vote.Rank)
		}

		voterIDs = append(voterIDs, vote.VoterID)
	}

	// One person one vote
	// First check that all votes come from the same voter
	valid := utils.AllEqual[string](voterIDs)
	if !valid {
		return false, fmt.Errorf("All votes must come from the same voter")
	}

	// Next check to make sure the person hasn't voted previously
	_, err := sm.Queries.GetVoterByIdIfVoted(context.Background(), db.GetVoterByIdIfVotedParams{ID: votes[0].VoterID, ElectionID: votes[0].ElectionID})
	if err == nil {
		return false, fmt.Errorf("voter has already cast a vote")
	}
	if err != sql.ErrNoRows {
		return false, err
	}
	return true, nil
}

func (sm *RankedChoice) ValidateVoter(voter Voter) (bool, error) {
	return ValidateVoterExists(sm.Queries, voter)
}
