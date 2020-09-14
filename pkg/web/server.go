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
	r.HandleFunc(baseUrlPattern, updateAndDisplay).Methods("GET")
	r.HandleFunc(baseUrlPattern+"/results", getResult).Methods("GET")
}

func castVote(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var vote protocol.Vote
	if err := decoder.Decode(&vote); err != nil {
		log.Fatal(err)
	}
	log.WithFields(log.Fields{"ballot": vote}).Info("Ballot received")
	election.Cast(vote)
}

func updateAndDisplay(w http.ResponseWriter, r *http.Request) {
	election.Tally()
	election.Display()
}

func getResult(w http.ResponseWriter, r *http.Request) {
	winner, total := election.Result()
	log.WithFields(log.Fields{"winner": winner, "total": total}).Info("Election results computed and returned:")
	msg := fmt.Sprintf("%s won with %g points", winner, total)
	w.Write([]byte(msg))
}
