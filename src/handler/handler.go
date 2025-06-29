package handler

import (
	"log"

	p "github.com/hokurikucar/PeriodTweetTrigger/src/post"
	t "github.com/hokurikucar/PeriodTweetTrigger/src/tweet"
)

// Handler Lambda関数のtrigger
func Handler() {
	po := p.NewPostObject()
	fw := p.NewPostFetchWorker()
	if err := fw.FetchPosts(po); err != nil {
		log.Fatal(err)
	}
	if err := t.Tweet(po.Title, po.URL, po.Tags); err != nil {
		log.Fatal(err)
	}
	log.Println("Completed to tweet!")
}
