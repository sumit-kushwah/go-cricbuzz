package cmd

import (
	"github.com/gocolly/colly"
)

type Series struct {
	Name      string
	Start_end string
}

func GetSeries(n int64, format string) []Series {
	series := []Series{}
	c := colly.NewCollector()

	c.OnHTML(".cb-sch-lst-itm", func(e *colly.HTMLElement) {
		item := Series{}
		item.Name = e.ChildText("span")
		item.Start_end = e.ChildText("div")
		series = append(series, item)
	})

	c.Visit("https://www.cricbuzz.com/cricket-schedule/series")
	c.Wait()
	return series[:n]
}
