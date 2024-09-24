package web

import (
	"fmt"
	"net/http"
	"time"

	"git.sr.ht/~bmaccini/go-vote/pkg/protocol"
	. "git.sr.ht/~bmaccini/go-vote/pkg/utils"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	Server   *http.Server
	Router   *chi.Mux
	Election protocol.Protocol
}

func CreateNewServer(election protocol.Protocol, port string) *Server {
	s := &Server{
		Election: election,
	}
	s.Router = chi.NewRouter()
	server := http.Server{
		Addr:              port,
		ReadHeaderTimeout: 3 * time.Second,
		Handler:           s.Router,
	}
	s.Server = &server
	return s
}

// Given a port start a web server on the port
func Init(p string, e protocol.Protocol) error {
	s := CreateNewServer(e, fmt.Sprintf(":%s", p))
	s.RegisterHandlers()

	// Launch the server on the specified port
	Logger.Info("Election server is live.", "electionId", e.GetID(), "port", p)
	err := s.Server.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) RegisterHandlers() {
	base := "/" + s.Election.GetID()

	// Add middleware
	s.Router.Use(middleware.Logger)
	s.Router.Use(middleware.RequestID)
	s.Router.Use(middleware.RealIP)
	s.Router.Use(middleware.Heartbeat(base + "/healthcheck"))
	s.Router.Use(middleware.Recoverer)

	// Voting
	voteRouter := chi.NewRouter()
	voteRouter.Post("/", s.CastVote)

	// ElectionResults
	electionResultsRouter := chi.NewRouter()
	electionResultsRouter.Get("/{electionId}", s.GetElectionResults)

	// Mount to the base router
	s.Router.Mount("/vote", voteRouter)
	s.Router.Mount("/results", electionResultsRouter)
}
