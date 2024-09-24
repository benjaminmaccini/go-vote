package cmd

import (
	"os"

	. "git.sr.ht/~bmaccini/go-vote/pkg/utils"

	"github.com/urfave/cli/v2"
)

var (
	dbName   string
	logLevel string
	port     string
)

var app = &cli.App{
	Name:  "go-vote",
	Usage: "A CLI app that handles elections",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:        "port",
			Aliases:     []string{"p"},
			Value:       "1337",
			Usage:       "port to host the election server on",
			Destination: &port,
			EnvVars:     []string{"GO_VOTE_PORT"},
		},
		&cli.StringFlag{
			Name:        "dbName",
			Aliases:     []string{"d"},
			Value:       "go-vote.sqlite",
			Usage:       "the database for the election",
			Destination: &dbName,
			EnvVars:     []string{"DATABASE_URL"},
		},
		&cli.StringFlag{
			Name:        "level",
			Value:       "INFO",
			Usage:       "The logging level",
			Destination: &logLevel,
			EnvVars:     []string{"LOG_LEVEL"},
		},
	},
	Commands: []*cli.Command{serveCmd},
}

func Execute() {
	InitLogger(logLevel)
	if err := app.Run(os.Args); err != nil {
		Logger.Fatal("", err)
	}
}
