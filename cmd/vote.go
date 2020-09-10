package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var voteCmd = &cobra.Command{
	Use:   "vote",
	Short: "Cast a ballot for an on-going election",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("vote called")
	},
}

func init() {
	rootCmd.AddCommand(voteCmd)

	voteCmd.Flags().StringP("ballot", "", "", "A json string representing a ballot")
	voteCmd.Flags().StringP("url", "", "", "The url of the election")
}
