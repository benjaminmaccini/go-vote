package web

import (
	"git.sr.ht/~bmaccini/go-vote/pkg/protocol"
)

type VoteRequest struct {
	ElectionId string        `json:"election_id"`
	Vote       protocol.Vote `json:"vote"`
}

type VoteResponse struct{}

type ElectionResultsRequest struct {
	ElectionId string `json:"electionId"`
}

type ElectionResultsRepsonse struct {
	ElectionResults map[string]float64 `json:"results"`
}
