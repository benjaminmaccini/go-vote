package cmd

import (
	"errors"
	"fmt"

	"git.sr.ht/~bmaccini/go-vote/pkg/db"
	"git.sr.ht/~bmaccini/go-vote/pkg/protocol"
	. "git.sr.ht/~bmaccini/go-vote/pkg/utils"
	"git.sr.ht/~bmaccini/go-vote/pkg/web"

	"github.com/urfave/cli/v2"
)

var protocolName string

var serveCmd = &cli.Command{
	Name:  "serve",
	Usage: "Start a server for an election",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:        "protocol",
			Usage:       "The election protocol to use",
			Destination: &protocolName,
		},
		&cli.StringSliceFlag{
			Name:  "candidate",
			Usage: "The name of a candidate for the election. Can be used multiple times",
		},
	},
	Action: func(ctx *cli.Context) error {
		candidates := []protocol.Candidate{}

		for _, name := range ctx.StringSlice("candidate") {
			candidates = append(candidates, protocol.Candidate{ID: GetStringUUIDv7(), Name: name})
		}

		if len(candidates) <= 1 {
			err := errors.New("Elections must have at least two candidates")
			Logger.Fatal("", err)
		}

		election, exists := protocol.ProtocolCommandMap[protocolName]
		if !exists {
			err := fmt.Errorf("Protocol %s does not exist", protocolName)
			Logger.Fatal("", err)
		}

		// Give access to the database for the election
		queries := db.InitDB(dbName)
		err := election.Init(candidates, queries)
		if err != nil {
			Logger.Fatal("Failed to initialize election", "error", err.Error())
		}
		err = web.Init(port, election)
		return err
	},
}
