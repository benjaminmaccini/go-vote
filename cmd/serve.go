package cmd

import (
	"fmt"
	"git.sr.ht/~bmaccini/go-vote/pkg/protocol"
	"git.sr.ht/~bmaccini/go-vote/pkg/web"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var protocolName string
var candidateNames []string

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Host a specific election protocol",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("serve called")

		candidates := []protocol.Candidate{}

		for _, name := range candidateNames {
			candidates = append(candidates, protocol.Candidate{Name: name})
		}

		election, exists := protocol.ProtocolCommandMap[protocolName]
		if !exists {
			fmt.Printf("Protocol %s does not exist", protocolName)
		}

		election.Init(candidates)
		web.Init(viper.GetString("port"), election)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.Flags().StringVar(&protocolName, "protocol", "", "The election protocol to use")
	serveCmd.Flags().StringSliceVar(&candidateNames, "candidates", []string{}, "A list of candidates for the election")
}
