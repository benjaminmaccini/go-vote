package protocol

import (
	"context"
	"testing"

	"git.sr.ht/~bmaccini/go-vote/pkg/db"
	. "git.sr.ht/~bmaccini/go-vote/pkg/utils"

	"github.com/stretchr/testify/assert"
)

// Happy path
func TestSimpleMajority(t *testing.T) {
	teardown := SetupTeardown(t)
	defer teardown(t)

	// Set up the election
	q := db.InitDB(TestDBPath)
	alice := Candidate{ID: GetStringUUIDv7(), Name: "alice"}
	bob := Candidate{ID: GetStringUUIDv7(), Name: "bob"}
	candidates := []Candidate{
		alice,
		bob,
	}
	election := new(SimpleMajority)
	err := election.Init(candidates, q)
	assert.Equal(t, err, nil, "")

	q.CreateVoter(context.Background(), db.CreateVoterParams{ID: "one"})
	q.CreateVoter(context.Background(), db.CreateVoterParams{ID: "two"})
	q.CreateVoter(context.Background(), db.CreateVoterParams{ID: "three"})
	err = election.Cast([]Vote{{
		ID:          GetStringUUIDv7(),
		CandidateID: alice.ID,
		VoterID:     "one",
		Rank:        1,
		ElectionID:  election.ID,
	}})
	assert.Equal(t, err, nil, "")
	err = election.Cast([]Vote{{
		ID:          GetStringUUIDv7(),
		CandidateID: alice.ID,
		VoterID:     "two",
		Rank:        1,
		ElectionID:  election.ID,
	}})
	assert.Equal(t, err, nil, "")
	err = election.Cast([]Vote{{
		ID:          GetStringUUIDv7(),
		CandidateID: bob.ID,
		VoterID:     "three",
		Rank:        1,
		ElectionID:  election.ID,
	}})
	assert.Equal(t, err, nil, "")

	// Test Result()
	results, err := election.Results()
	assert.Equal(t, err, nil, "")

	assert.Equal(t, 2., results[alice.ID], "")
	assert.Equal(t, 1., results[bob.ID], "")
}

func TestSimpleMajorityValidate(t *testing.T) {
	teardown := SetupTeardown(t)
	defer teardown(t)

	// Set up the election
	q := db.InitDB(TestDBPath)
	voter := Voter{ID: GetStringUUIDv7()}
	q.CreateVoter(context.Background(), db.CreateVoterParams{ID: voter.ID})
	alice := Candidate{ID: GetStringUUIDv7(), Name: "alice"}
	bob := Candidate{ID: GetStringUUIDv7(), Name: "bob"}
	candidates := []Candidate{
		alice,
		bob,
	}
	election := new(SimpleMajority)
	err := election.Init(candidates, q)
	assert.Equal(t, err, nil, "")

	// Multiple vote in the payload
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
			VoterID:     GetStringUUIDv7(),
			Rank:        1,
			ElectionID:  election.ID,
		},
	})
	assert.Equal(t, valid, false, "")

	// Candidate doesn't exist
	valid, err = election.ValidateVote([]Vote{{
		ID:          GetStringUUIDv7(),
		CandidateID: GetStringUUIDv7(),
		VoterID:     voter.ID,
		Rank:        1,
		ElectionID:  election.ID,
	}})
	assert.Equal(t, valid, false, "")

	// Wrong election
	valid, err = election.ValidateVote([]Vote{{
		ID:          GetStringUUIDv7(),
		CandidateID: alice.ID,
		VoterID:     voter.ID,
		Rank:        1,
		ElectionID:  GetStringUUIDv7(),
	}})
	assert.Equal(t, valid, false, "")

	// Multiple votes, one person
	votes := []Vote{{
		ID:          GetStringUUIDv7(),
		CandidateID: alice.ID,
		VoterID:     voter.ID,
		Rank:        1,
		ElectionID:  election.ID,
	}}
	valid, err = election.ValidateVote(votes)
	assert.Equal(t, valid, true, "")
	err = election.Cast(votes)
	assert.Equal(t, err, nil, "")
	valid, err = election.ValidateVote(votes)
	assert.Equal(t, valid, false, "")
}
