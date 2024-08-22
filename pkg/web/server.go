package web

import (
	"encoding/json"
	"fmt"
	"net/http"

	. "git.sr.ht/~bmaccini/go-vote/pkg/utils"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"git.sr.ht/~bmaccini/go-vote/pkg/protocol"
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

	InitLogger("INFO")

	// Launch the server on the specified port
	port := fmt.Sprintf(":%s", p)
	Logger.Info("Election server is live.", "electionId", protocol.E.GetId(), "port", port)
	Logger.Fatal("", http.ListenAndServe(port, s.Router))
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
		Logger.Error("", err)
		w.WriteHeader(400)
		return
	}
	valid := protocol.E.ValidateVote(vote)
	if !valid {
		Logger.Error("Invalid vote received", "ballot", vote)
		w.WriteHeader(406)
		return
	}
	protocol.E.Cast(vote)

	w.WriteHeader(200)
}

func getResult(w http.ResponseWriter, r *http.Request) {
	winners, total := protocol.E.Result()
	Logger.Info("Election results computed and returned:", "winners", winners, "total", total)
	msg := fmt.Sprintf("%s won with %f points", winners, total)
	w.WriteHeader(200)
	w.Write([]byte(msg))
}
