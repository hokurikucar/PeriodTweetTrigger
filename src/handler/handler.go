package handler

import (
	"log"

	q "github.com/hokurikucar/PeriodTweetTrigger/src/queryFetcher"
	t "github.com/hokurikucar/PeriodTweetTrigger/src/tweet"
)

// Handler Lambda関数のtrigger
func Handler() {
	a := q.NewArticleObject()
	if err := a.FetchArticles(); err != nil {
		log.Fatal("Fetching articles error: %+v", err)
	}

	if err := t.Tweet(a.Title, a.URL); err != nil {
		log.Fatal("Posting tweet error: %+v", err)
	}

	log.Println("Completed to tweet!")
}
