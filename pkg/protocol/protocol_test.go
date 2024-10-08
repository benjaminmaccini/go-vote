package protocol

import (
	"testing"

	. "git.sr.ht/~bmaccini/go-vote/pkg/utils"
)

// Test that the functions in the map are properly defined
func TestProtocolNameMap(t *testing.T) {
	for _, election := range ProtocolCommandMap {
		election.Init([]Candidate{Candidate{Name: "alice"}})
		election.Display()
	}
}

// TODO split up the test learn mocks
func TestSimpleMajority(t *testing.T) {
	alice := Candidate{Name: "alice"}
	bob := Candidate{Name: "bob"}
	candidates := []Candidate{
		alice,
		bob,
	}

	election := new(SimpleMajority)
	// Test Init()
	election.Init(candidates)

	// Test Cast()
	voteA := Vote{Candidate: alice, Value: 1.0}

	election.Cast(voteA)

	voteB := Vote{Candidate: bob, Value: 1.0}
	voteA2 := Vote{Candidate: alice, Value: 1.0}
	election.Cast(voteB)
	election.Cast(voteA2)

	// Test Tally()
	election.Tally()

	// Test Display()
	election.Display()

	voteB2 := Vote{Candidate: bob, Value: 1.0}
	voteB3 := Vote{Candidate: bob, Value: 1.0}
	election.Cast(voteB2)
	election.Cast(voteB3)

	// Test Result()
	winners, count := election.Result()

	election.Display()

	AssertEqual(t, 1, len(winners), "")
	AssertEqual(t, bob.Name, winners[0], "")
	AssertEqual(t, 3.0, count, "")
}
