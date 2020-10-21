package handler

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

const hokurikuCarURL = "https://hokurikucar.com/"
const articleSelectorPath = "article > div > h2 > a"

func Handler() {

	log.Println("started")
	// Webサイトへのリクエスト
	res, err := http.Get(hokurikuCarURL)
	if err != nil {
		log.Fatal("Http Request Error: %s", err)
	}
	defer res.Body.Close()

	log.Println(res)
	if res.StatusCode != 200 {
		log.Fatal("Response from website is not 200: %d %s", res.StatusCode, res.Status)
	}

	// HTMLドキュメントの取得
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal("Load HTML Document Error: %s", err)
	}
	log.Println(doc)
	doc.Find(articleSelectorPath).Each(func(i int, s *goquery.Selection) {
		// 取得できた記事１つ１つに対する処理
		title := s.Text()
		url, _ := s.Attr("href")
		log.Println(title)
		log.Println(url)
	})
}
