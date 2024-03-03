package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var matchType string
var team string

func getMatches(scoreUrl string, n int) {

}

func init() {
	matchCommand.Flags().StringVarP(&matchType, "type", "t", "live", "live, recent, upcoming")
	matchCommand.Flags().StringVar(&team, "team", "", "true or false")
	matchCommand.Flags().IntVarP(&count, "number", "n", 5, "true or false")
	rootCmd.AddCommand(matchCommand)
}

var matchCommand = &cobra.Command{
	Use:   "match",
	Short: "Match information",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		var scoreUrl string
		if matchType == "live" {
			scoreUrl = "https://www.cricbuzz.com/cricket-match/live-scores"
		} else if matchType == "recent" {
			scoreUrl = "https://www.cricbuzz.com/cricket-match/live-scores/recent-matches"
		} else if matchType == "upcoming" {
			scoreUrl = "https://www.cricbuzz.com/cricket-match/live-scores/upcoming-matches"
		} else {
			fmt.Println("Match type is wrong.")
		}

		if len(team) > 0 {

		} else {

		}

	},
}
