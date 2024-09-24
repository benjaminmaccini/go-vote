package protocol

import (
	"testing"

	"git.sr.ht/~bmaccini/go-vote/pkg/db"
	. "git.sr.ht/~bmaccini/go-vote/pkg/utils"

	"github.com/stretchr/testify/assert"
)

func TestSimpleMajority(t *testing.T) {
	teardown := SetupTeardown(t)
	defer teardown(t)

	// Set up the election
	q := db.InitDB(TestDBPath)
	voter := Voter{ID: GetStringUUIDv7()}
	alice := Candidate{ID: GetStringUUIDv7(), Name: "alice"}
	bob := Candidate{ID: GetStringUUIDv7(), Name: "bob"}
	candidates := []Candidate{
		alice,
		bob,
	}
	election := new(SimpleMajority)
	err := election.Init(candidates, q)
	assert.Equal(t, err, nil, "")

	err = election.Cast(Vote{
		ID:          GetStringUUIDv7(),
		CandidateID: alice.ID,
		VoterID:     voter.ID,
		Rank:        1,
	})
	assert.Equal(t, err, nil, "")
	err = election.Cast(Vote{
		ID:          GetStringUUIDv7(),
		CandidateID: alice.ID,
		VoterID:     voter.ID,
		Rank:        1,
	})
	assert.Equal(t, err, nil, "")
	err = election.Cast(Vote{
		ID:          GetStringUUIDv7(),
		CandidateID: bob.ID,
		VoterID:     voter.ID,
		Rank:        1,
	})
	assert.Equal(t, err, nil, "")

	// Test Result()
	results, err := election.Results()
	assert.Equal(t, err, nil, "")

	assert.Equal(t, 2., results[alice.ID], "")
	assert.Equal(t, 1., results[bob.ID], "")
}
