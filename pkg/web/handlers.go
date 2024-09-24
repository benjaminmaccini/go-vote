package web

import (
	"encoding/json"
	"net/http"

	"git.sr.ht/~bmaccini/go-vote/pkg/protocol"
	. "git.sr.ht/~bmaccini/go-vote/pkg/utils"
	"github.com/go-chi/chi"
)

func (s *Server) CastVote(w http.ResponseWriter, r *http.Request) {
	var vote protocol.Vote
	err := json.NewDecoder(r.Body).Decode(&vote)
	if err != nil {
		Logger.Error("", err)
		w.WriteHeader(400)
		return
	}

	valid, err := s.Election.ValidateVote(vote)
	if !valid && err != nil {
		Logger.Error("Invalid vote received", "ballot", vote)
		w.WriteHeader(406)
		return
	}

	err = s.Election.Cast(vote)
	if err != nil {
		Logger.Error("Error submitting the vote", "error", err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(200)
	return
}

func (s *Server) GetElectionResults(w http.ResponseWriter, r *http.Request) {
	_ = chi.URLParam(r, "electionId")
	results, err := s.Election.Results()
	if err != nil {
		Logger.Error("Error fetching the election results", "error", err)
		w.WriteHeader(500)
		return
	}

	err = json.NewEncoder(w).Encode(results)
	if err != nil {
		Logger.Error("", err)
		Logger.Error("Error getting election results")
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(200)
	return
}
