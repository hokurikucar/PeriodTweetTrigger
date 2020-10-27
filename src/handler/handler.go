package handler

import (
	"log"

	p "github.com/hokurikucar/PeriodTweetTrigger/src/post"
	t "github.com/hokurikucar/PeriodTweetTrigger/src/tweet"
)

// Handler Lambda関数のtrigger
func Handler() {
	a := p.NewArticleObject()
	if err := a.FetchArticles(); err != nil {
		log.Fatal("Fetching articles error: %+v", err)
	}

	if err := t.Tweet(a.Title, a.URL); err != nil {
		log.Fatal("Posting tweet error: %+v", err)
	}

	log.Println("Completed to tweet!")
}
