package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var protocol string

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Host a specific election protocol",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("serve called")
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.Flags().StringVar(&protocol, "protocol", "", "The election protocol to use")
}
