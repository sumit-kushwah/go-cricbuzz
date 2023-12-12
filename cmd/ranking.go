package cmd

import (
	"fmt"
	"go-cricbuzz/common"

	"github.com/gocolly/colly"
	"github.com/spf13/cobra"
)

type TeamRanking struct {
	Name     string
	Position string
	Rating   string
	Points   string
}

type PlayerRanking struct {
	Name     string
	Position string
	Country  string
	Rating   string
}

func GetTeamsRankings(format string, n int) []TeamRanking {
	teamRankings := []TeamRanking{}
	c := colly.NewCollector()
	c.OnHTML(fmt.Sprintf("div[ng-show=\"'teams-%ss' == act_rank_format\"]", format), func(e *colly.HTMLElement) {
		e.ForEach("div.cb-col.cb-col-100.cb-font-14.text-center", func(i int, h *colly.HTMLElement) {
			position_rating := []string{}
			h.ForEach("div.cb-col.cb-col-14", func(i int, g *colly.HTMLElement) {
				position_rating = append(position_rating, g.Text)
			})
			points := ""
			if len(position_rating) > 1 {
				points = position_rating[1]
			}
			rank := TeamRanking{
				Name:     h.ChildText("div.cb-col.cb-col-50.text-left"),
				Position: h.ChildText("div.cb-col.cb-col-20"),
				Rating:   position_rating[0],
				Points:   points,
			}
			teamRankings = append(teamRankings, rank)
		})
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Printf("Something went wrong\nStatus code from cricbuzz: %d", r.StatusCode)
	})

	c.Visit("https://www.cricbuzz.com/cricket-stats/icc-rankings/men/teams")
	c.Wait()
	if len(teamRankings) > n {
		return teamRankings[0 : n+1]
	}
	return teamRankings
}

func GetPlayerRankings(playerType string, format string, n int, isMen bool) []PlayerRanking {
	playerRankings := []PlayerRanking{}
	c := colly.NewCollector()
	c.OnHTML(fmt.Sprintf("div[ng-show=\"'%s-%s' == act_rank_format\"]", playerType, format), func(e *colly.HTMLElement) {
		e.ForEach("div.cb-col.cb-col-100.cb-font-14.text-center", func(i int, h *colly.HTMLElement) {
			rank := PlayerRanking{
				Name:     h.ChildText("a.text-hvr-underline.text-bold.cb-font-16"),
				Position: h.ChildText("div.cb-col.cb-col-16"),
				Rating:   h.ChildText("div.cb-col.cb-col-17"),
			}
			playerRankings = append(playerRankings, rank)
		})
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Printf("Something went wrong\nStatus code from cricbuzz: %d", r.StatusCode)
	})

	var endpoint string
	if playerType == "batsmen" {
		endpoint = "batting"
	} else if playerType == "bowlers" {
		endpoint = "bowling"
	} else if playerType == "allrounders" {
		endpoint = "all-rounder"
	}

	gender := "women"
	if isMen {
		gender = "men"
	}

	c.Visit(fmt.Sprintf("https://www.cricbuzz.com/cricket-stats/icc-rankings/%s/%s", gender, endpoint))
	c.Wait()
	if len(playerRankings) > n {
		return playerRankings[0 : n+1]
	}
	return playerRankings
}

func PrintPlayerRankings(rankings []PlayerRanking) {
	for _, item := range rankings {
		fmt.Printf("%s %s %s\n",
			common.FormatInNLength(item.Position, 9),
			common.FormatInNLength(item.Name, 10),
			common.FormatInNLength(item.Rating, 8),
		)
	}
}

var format string
var isMen bool
var playerType string
var count int

func init() {
	rankingCommand.Flags().StringVarP(&format, "format", "f", "tests", "Cricket Format")
	rankingCommand.Flags().BoolVar(&isMen, "men", true, "For Men")
	rankingCommand.Flags().StringVarP(&playerType, "type", "t", "batsmen", "Cricket Format")
	rankingCommand.Flags().IntVarP(&count, "number", "n", 5, "Number of records")
	rootCmd.AddCommand(rankingCommand)
}

var rankingCommand = &cobra.Command{
	Use:   "ranking",
	Short: "List Icc rankings",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		// rankings := GetPlayerRankings(playerType, format, count, isMen)
		// PrintPlayerRankings(rankings)
	},
}
