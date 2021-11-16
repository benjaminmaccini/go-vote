package web

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/benjaminmaccini/go-vote/pkg/protocol"
)

// From go-chi documentation. Executes a request
func executeRequest(req *http.Request, s *Server) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	s.Router.ServeHTTP(rr, req)

	return rr
}

func createTestElection(candidateNames []string, protocolName string) protocol.Protocol {
	candidates := []protocol.Candidate{}

	for _, name := range candidateNames {
		candidates = append(candidates, protocol.Candidate{Name: name})
	}

	election, exists := protocol.ProtocolCommandMap[protocolName]
	if !exists {
		fmt.Printf("Protocol %s does not exist", protocolName)
	}

	election.Init(candidates)

	election.SetId("testId")

	election.Display()

	return election
}

func TestAlive(t *testing.T) {
	s := CreateNewServer("testId")
	s.RegisterHandlers()

	req, _ := http.NewRequest("GET", "/testId/healthcheck", nil)
	resp := executeRequest(req, s)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, ".", resp.Body.String())
}

func TestCastVote(t *testing.T) {
	s := CreateNewServer("testId")
	s.RegisterHandlers()

	e := createTestElection([]string{"Alice"}, "simpleMajority")

	protocol.E = e

	vote := protocol.Vote{
		Value:     1.0,
		Candidate: protocol.Candidate{Name: "Alice"},
	}

	j, err := json.Marshal(vote)

	if err != nil {
		t.Error(err)
	}

	req, _ := http.NewRequest("POST", "/testId/vote", bytes.NewBuffer(j))
	resp := executeRequest(req, s)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, "", resp.Body.String())
}

// Cast a vote for an non-existent candidate
func TestCastVoteFail(t *testing.T) {
	s := CreateNewServer("testId")
	s.RegisterHandlers()

	e := createTestElection([]string{"Alice"}, "simpleMajority")

	protocol.E = e

	vote := protocol.Vote{
		Value:     1.0,
		Candidate: protocol.Candidate{Name: "Bob"},
	}

	j, err := json.Marshal(vote)

	if err != nil {
		t.Error(err)
	}

	req, _ := http.NewRequest("POST", "/testId/vote", bytes.NewBuffer(j))
	resp := executeRequest(req, s)

	assert.Equal(t, http.StatusNotAcceptable, resp.Code)
}

func TestResults(t *testing.T) {
	s := CreateNewServer("testId")
	s.RegisterHandlers()

	e := createTestElection([]string{"Alice", "Bob"}, "simpleMajority")

	protocol.E = e

	votes := []protocol.Vote{
		{
			Value:     1.0,
			Candidate: protocol.Candidate{Name: "Alice"},
		},
		{
			Value:     1.0,
			Candidate: protocol.Candidate{Name: "Bob"},
		},
		{
			Value:     1.0,
			Candidate: protocol.Candidate{Name: "Alice"},
		},
		{
			Value:     1.0,
			Candidate: protocol.Candidate{Name: "Bob"},
		},
		{
			Value:     1.0,
			Candidate: protocol.Candidate{Name: "Alice"},
		},
	}

	// Cast all the votes
	for _, v := range votes {
		j, err := json.Marshal(v)

		if err != nil {
			t.Error(err)
			t.Fail()
		}

		req, _ := http.NewRequest("POST", "/testId/vote", bytes.NewBuffer(j))
		_ = executeRequest(req, s)
	}

	// Find the results
	expected := fmt.Sprintf("%s won with %f points", "[Alice]", 3.0)
	req, _ := http.NewRequest("GET", "/testId/results", nil)
	resp := executeRequest(req, s)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, expected, resp.Body.String())
}
