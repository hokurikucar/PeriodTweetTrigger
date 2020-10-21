package handler

import (
	"log"

	q "github.com/hokurikucar/PeriodTweetTrigger/src/queryFetcher"
)

// Handler Lambda関数のtrigger
func Handler() {
	log.Println("処理開始")
	a := q.NewArticleObject()
	a.FetchArticles()
}
