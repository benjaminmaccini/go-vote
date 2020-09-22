package web

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/benjaminmaccini/go-vote/pkg/protocol"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

var election protocol.Protocol

// Given a port start a web server on the port
func Init(p string, e protocol.Protocol) {
	election = e
	r := mux.NewRouter()
	registerHandlers(r, election.GetId())

	// Launch the server on the specified port
	port := fmt.Sprintf(":%s", p)
	log.WithFields(log.Fields{"electionId": election.GetId(), "port": port}).Info("Election server is live.")
	log.Fatal(http.ListenAndServe(port, r))
}

func registerHandlers(r *mux.Router, electionId string) {
	baseUrlPattern := fmt.Sprintf("/%s", electionId)
	r.HandleFunc(baseUrlPattern, castVote).Methods("POST")
	r.HandleFunc(baseUrlPattern+"/results", getResult).Methods("GET")
}

func castVote(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var vote protocol.Vote
	if err := decoder.Decode(&vote); err != nil {
		log.Error(err)
		return
	}
	valid := election.ValidateVote(vote)
	if !valid {
		log.WithFields(log.Fields{"ballot": vote}).Error("Invalid vote received")
		return
	}
	election.Cast(vote)
}

func getResult(w http.ResponseWriter, r *http.Request) {
	winners, total := election.Result()
	log.WithFields(log.Fields{"winners": winners, "total": total}).Info("Election results computed and returned:")
	msg := fmt.Sprintf("%s won with %f points", winners, total)
	w.Write([]byte(msg))
}
