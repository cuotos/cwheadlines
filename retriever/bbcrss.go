package retriever

import (
	"github.com/mmcdole/gofeed"
	"math/rand"
)

type BbcRss struct{}

func (BbcRss) GetSolution() (string, error) {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL("http://feeds.bbci.co.uk/news/rss.xml")
	if err != nil {
		return "", err
	}

	item := feed.Items[rand.Intn(len(feed.Items))]

	return item.Title, err
}



