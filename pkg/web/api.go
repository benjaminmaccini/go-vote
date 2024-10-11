package web

import (
	"git.sr.ht/~bmaccini/go-vote/pkg/protocol"
)

type VoteRequest struct {
	Voter protocol.Voter  `json:"voter"`
	Votes []protocol.Vote `json:"votes"`
}

type VoteResponse struct{}

type ElectionResultsRequest struct {
	ElectionId string `json:"electionId"`
}

type ElectionResultsRepsonse struct {
	ElectionResults map[string]float64 `json:"results"`
}
