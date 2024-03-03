package cmd

import (
	"fmt"

	"github.com/gocolly/colly"
)

type Player struct {
	name string
}

func GetPlayerDetails(searchTerm string) {
	c := colly.NewCollector()
	c.OnHTML("span[\"ng-bind=result.title\"]", func(e *colly.HTMLElement) {
		
		fmt.Println(e.Text)
	})

	c.Visit(fmt.Sprintf("https://www.cricbuzz.com/search?q=%s&tab=player", searchTerm))

}
