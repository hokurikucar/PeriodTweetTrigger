package handler

import (
	"log"

	q "github.com/hokurikucar/PeriodTweetTrigger/src/queryFetcher"
)

// Handler Lambda関数のtrigger
func Handler() {
	log.Println("処理開始")

	log.Println("記事の取得処理を開始")
	a := q.NewArticleObject()
	if err := a.FetchArticles(); err != nil {
		log.Fatal(err)
	}
	log.Println("%+v", a)
	log.Println("記事の取得処理を終了")
}
