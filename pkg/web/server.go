package web

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	log "github.com/sirupsen/logrus"

	"github.com/benjaminmaccini/go-vote/pkg/protocol"
)

type Server struct {
	Router     *chi.Mux
	ElectionId string
}

func CreateNewServer(eid string) *Server {
	s := &Server{ElectionId: eid}
	s.Router = chi.NewRouter()
	return s
}

// Given a port start a web server on the port
func Init(p string, e protocol.Protocol) {
	// IMPORTANT: Set the current election based on the CLI
	protocol.E = e

	s := CreateNewServer(protocol.E.GetId())
	s.RegisterHandlers()

	// Launch the server on the specified port
	port := fmt.Sprintf(":%s", p)
	log.WithFields(log.Fields{"electionId": protocol.E.GetId(), "port": port}).Info("Election server is live.")
	log.Fatal(http.ListenAndServe(port, s.Router))
}

func (s *Server) RegisterHandlers() {
	base := "/" + s.ElectionId

	// Add middleware
	s.Router.Use(middleware.Logger)
	s.Router.Use(middleware.Heartbeat(base + "/healthcheck"))
	s.Router.Use(middleware.Recoverer)

	// Add routes
	s.Router.Post(base+"/vote", castVote)
	s.Router.Get(base+"/results", getResult)
}

func castVote(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var vote protocol.Vote
	if err := decoder.Decode(&vote); err != nil {
		log.Error(err)
		w.WriteHeader(400)
		return
	}
	valid := protocol.E.ValidateVote(vote)
	if !valid {
		log.WithFields(log.Fields{"ballot": vote}).Error("Invalid vote received")
		w.WriteHeader(406)
		return
	}
	protocol.E.Cast(vote)

	w.WriteHeader(200)
}

func getResult(w http.ResponseWriter, r *http.Request) {
	winners, total := protocol.E.Result()
	log.WithFields(log.Fields{"winners": winners, "total": total}).Info("Election results computed and returned:")
	msg := fmt.Sprintf("%s won with %f points", winners, total)
	w.WriteHeader(200)
	w.Write([]byte(msg))
}
