package protocol

import (
	"context"
	"testing"

	"git.sr.ht/~bmaccini/go-vote/pkg/db"
	. "git.sr.ht/~bmaccini/go-vote/pkg/utils"

	"github.com/stretchr/testify/assert"
)

// Happy path get results
func TestRankedChoice(t *testing.T) {
	teardown := SetupTeardown(t)
	defer teardown(t)

	// Set up the election
	q := db.InitDB(TestDBPath)
	alice := Candidate{ID: GetStringUUIDv7(), Name: "alice"}
	bob := Candidate{ID: GetStringUUIDv7(), Name: "bob"}
	charlie := Candidate{ID: GetStringUUIDv7(), Name: "charlie"}
	candidates := []Candidate{
		alice,
		bob,
		charlie,
	}
	election := new(RankedChoice)
	err := election.Init(candidates, q)
	assert.Equal(t, err, nil, "")

	voter := Voter{ID: GetStringUUIDv7()}
	q.CreateVoter(context.Background(), db.CreateVoterParams{ID: voter.ID})

	err = election.Cast([]Vote{
		{
			ID:          GetStringUUIDv7(),
			CandidateID: alice.ID,
			VoterID:     voter.ID,
			Rank:        1,
			ElectionID:  election.ID,
		},
		{
			ID:          GetStringUUIDv7(),
			CandidateID: bob.ID,
			VoterID:     voter.ID,
			Rank:        2,
			ElectionID:  election.ID,
		},
		{
			ID:          GetStringUUIDv7(),
			CandidateID: charlie.ID,
			VoterID:     voter.ID,
			Rank:        3,
			ElectionID:  election.ID,
		},
	})
	assert.Equal(t, err, nil, "")

	// Reinit the voter
	voter = Voter{ID: GetStringUUIDv7()}
	q.CreateVoter(context.Background(), db.CreateVoterParams{ID: voter.ID})
	err = election.Cast([]Vote{
		{
			ID:          GetStringUUIDv7(),
			CandidateID: alice.ID,
			VoterID:     voter.ID,
			Rank:        2,
			ElectionID:  election.ID,
		},
		{
			ID:          GetStringUUIDv7(),
			CandidateID: bob.ID,
			VoterID:     voter.ID,
			Rank:        1,
			ElectionID:  election.ID,
		},
		{
			ID:          GetStringUUIDv7(),
			CandidateID: charlie.ID,
			VoterID:     voter.ID,
			Rank:        3,
			ElectionID:  election.ID,
		},
	})
	assert.Equal(t, err, nil, "")

	// Reinit the voter
	voter = Voter{ID: GetStringUUIDv7()}
	q.CreateVoter(context.Background(), db.CreateVoterParams{ID: voter.ID})
	err = election.Cast([]Vote{
		{
			ID:          GetStringUUIDv7(),
			CandidateID: alice.ID,
			VoterID:     voter.ID,
			Rank:        2,
			ElectionID:  election.ID,
		},
		{
			ID:          GetStringUUIDv7(),
			CandidateID: bob.ID,
			VoterID:     voter.ID,
			Rank:        3,
			ElectionID:  election.ID,
		},
		{
			ID:          GetStringUUIDv7(),
			CandidateID: charlie.ID,
			VoterID:     voter.ID,
			Rank:        1,
			ElectionID:  election.ID,
		},
	})
	assert.Equal(t, err, nil, "")

	// Test Result()
	results, err := election.Results()
	assert.Equal(t, err, nil, "")

	assert.Equal(t, 5., results[alice.ID], "")
	assert.Equal(t, 6., results[bob.ID], "")
	assert.Equal(t, 7., results[charlie.ID], "")
}

func TestRankedChoiceBadVotes(t *testing.T) {
	teardown := SetupTeardown(t)
	defer teardown(t)

	// Set up the election
	q := db.InitDB(TestDBPath)
	alice := Candidate{ID: GetStringUUIDv7(), Name: "alice"}
	bob := Candidate{ID: GetStringUUIDv7(), Name: "bob"}
	charlie := Candidate{ID: GetStringUUIDv7(), Name: "charlie"}
	candidates := []Candidate{
		alice,
		bob,
		charlie,
	}
	election := new(RankedChoice)
	err := election.Init(candidates, q)
	assert.Equal(t, err, nil, "")

	// All candidates accounted for
	voter := Voter{ID: GetStringUUIDv7()}
	q.CreateVoter(context.Background(), db.CreateVoterParams{ID: voter.ID})
	valid, err := election.ValidateVote([]Vote{
		{
			ID:          GetStringUUIDv7(),
			CandidateID: bob.ID,
			VoterID:     voter.ID,
			Rank:        1,
			ElectionID:  election.ID,
		},
		{
			ID:          GetStringUUIDv7(),
			CandidateID: bob.ID,
			VoterID:     voter.ID,
			Rank:        2,
			ElectionID:  election.ID,
		},
		{
			ID:          GetStringUUIDv7(),
			CandidateID: bob.ID,
			VoterID:     voter.ID,
			Rank:        3,
			ElectionID:  election.ID,
		},
	})
	assert.Equal(t, valid, false, "")

	// Not too many votes
	voter = Voter{ID: GetStringUUIDv7()}
	q.CreateVoter(context.Background(), db.CreateVoterParams{ID: voter.ID})
	valid, err = election.ValidateVote([]Vote{
		{
			ID:          GetStringUUIDv7(),
			CandidateID: alice.ID,
			VoterID:     voter.ID,
			Rank:        1,
			ElectionID:  election.ID,
		},
		{
			ID:          GetStringUUIDv7(),
			CandidateID: bob.ID,
			VoterID:     voter.ID,
			Rank:        2,
			ElectionID:  election.ID,
		},
		{
			ID:          GetStringUUIDv7(),
			CandidateID: charlie.ID,
			VoterID:     voter.ID,
			Rank:        3,
			ElectionID:  election.ID,
		},
		{
			ID:          GetStringUUIDv7(),
			CandidateID: charlie.ID,
			VoterID:     voter.ID,
			Rank:        3,
			ElectionID:  election.ID,
		},
	})
	assert.Equal(t, valid, false, "")

	// Not too little votes
	voter = Voter{ID: GetStringUUIDv7()}
	q.CreateVoter(context.Background(), db.CreateVoterParams{ID: voter.ID})
	valid, err = election.ValidateVote([]Vote{
		{
			ID:          GetStringUUIDv7(),
			CandidateID: alice.ID,
			VoterID:     voter.ID,
			Rank:        1,
			ElectionID:  election.ID,
		},
		{
			ID:          GetStringUUIDv7(),
			CandidateID: bob.ID,
			VoterID:     voter.ID,
			Rank:        2,
			ElectionID:  election.ID,
		},
	})
	assert.Equal(t, valid, false, "")

	// Ranks are complete based on number of candidates
	voter = Voter{ID: GetStringUUIDv7()}
	q.CreateVoter(context.Background(), db.CreateVoterParams{ID: voter.ID})
	valid, err = election.ValidateVote([]Vote{
		{
			ID:          GetStringUUIDv7(),
			CandidateID: alice.ID,
			VoterID:     voter.ID,
			Rank:        1,
			ElectionID:  election.ID,
		},
		{
			ID:          GetStringUUIDv7(),
			CandidateID: bob.ID,
			VoterID:     voter.ID,
			Rank:        100,
			ElectionID:  election.ID,
		},
		{
			ID:          GetStringUUIDv7(),
			CandidateID: charlie.ID,
			VoterID:     voter.ID,
			Rank:        3,
			ElectionID:  election.ID,
		},
	})
	assert.Equal(t, valid, false, "")

	// Ranks are not redundant
	voter = Voter{ID: GetStringUUIDv7()}
	q.CreateVoter(context.Background(), db.CreateVoterParams{ID: voter.ID})
	valid, err = election.ValidateVote([]Vote{
		{
			ID:          GetStringUUIDv7(),
			CandidateID: alice.ID,
			VoterID:     voter.ID,
			Rank:        1,
			ElectionID:  election.ID,
		},
		{
			ID:          GetStringUUIDv7(),
			CandidateID: bob.ID,
			VoterID:     voter.ID,
			Rank:        1,
			ElectionID:  election.ID,
		},
		{
			ID:          GetStringUUIDv7(),
			CandidateID: charlie.ID,
			VoterID:     voter.ID,
			Rank:        3,
			ElectionID:  election.ID,
		},
	})
	assert.Equal(t, valid, false, "")

	// Votes come from a different voter
	voter = Voter{ID: GetStringUUIDv7()}
	q.CreateVoter(context.Background(), db.CreateVoterParams{ID: voter.ID})
	valid, err = election.ValidateVote([]Vote{
		{
			ID:          GetStringUUIDv7(),
			CandidateID: alice.ID,
			VoterID:     voter.ID,
			Rank:        2,
			ElectionID:  election.ID,
		},
		{
			ID:          GetStringUUIDv7(),
			CandidateID: bob.ID,
			VoterID:     voter.ID,
			Rank:        3,
			ElectionID:  election.ID,
		},
		{
			ID:          GetStringUUIDv7(),
			CandidateID: charlie.ID,
			VoterID:     "DIFFERENT_VOTER",
			Rank:        1,
			ElectionID:  election.ID,
		},
	})
	assert.Equal(t, valid, false, "")

	// One person one vote
	voter = Voter{ID: GetStringUUIDv7()}
	q.CreateVoter(context.Background(), db.CreateVoterParams{ID: voter.ID})
	votes := []Vote{
		{
			ID:          GetStringUUIDv7(),
			CandidateID: alice.ID,
			VoterID:     voter.ID,
			Rank:        1,
			ElectionID:  election.ID,
		},
		{
			ID:          GetStringUUIDv7(),
			CandidateID: bob.ID,
			VoterID:     voter.ID,
			Rank:        2,
			ElectionID:  election.ID,
		},
		{
			ID:          GetStringUUIDv7(),
			CandidateID: charlie.ID,
			VoterID:     voter.ID,
			Rank:        3,
			ElectionID:  election.ID,
		},
	}
	valid, err = election.ValidateVote(votes)
	assert.Equal(t, valid, true, "")
	// We must cast the votes in order to fail
	err = election.Cast(votes)
	valid, err = election.ValidateVote(votes)
	assert.Equal(t, valid, false, "")
}
