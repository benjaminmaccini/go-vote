package cmd

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var baseUrl string
var ballot string
var electionId string

var voteCmd = &cobra.Command{
	Use:   "vote",
	Short: "Cast a ballot for an on-going election",
	Run: func(cmd *cobra.Command, args []string) {
		requestUrl := fmt.Sprintf("http://%s:%s/%s", baseUrl, viper.GetString("port"), electionId)
		data := []byte(ballot)

		req, err := http.NewRequest("POST", requestUrl, bytes.NewBuffer(data))
		if err != nil {
			log.WithFields(log.Fields{"err": err}).Error("Bad request")
			log.WithFields(log.Fields{
				"ballot":     ballot,
				"url":        requestUrl,
				"electionId": electionId,
			}).Fatal("Bad request")
		}

		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{Timeout: time.Second * 10}

		resp, err := client.Do(req)
		if err != nil {
			log.WithFields(log.Fields{"err": err}).Fatal("Bad request")
		}
		defer resp.Body.Close()
	},
}

func init() {
	rootCmd.AddCommand(voteCmd)

	voteCmd.Flags().StringVar(&ballot, "ballot", "", "A json string representing a ballot")
	voteCmd.Flags().StringVar(&baseUrl, "url", "", "The url of the election server")
	voteCmd.Flags().StringVar(&electionId, "electionId", "", "The id of the election")
}
