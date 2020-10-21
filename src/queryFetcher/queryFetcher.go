package queryFetcher

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

const hokurikuCarURL = "https://hokurikucar.com/"
const articleSelectorPath = "article > div > h2 > a"

// Article 記事の情報を格納するオブジェクト
type Article struct {
	Title string
	URL   string
}

// NewArticleObject 取得した記事情報を格納するオブジェクトを生成して返却する
func NewArticleObject() *Article {
	return &Article{}
}

// FetchArticles 記事のタイトルとURLを取得してオブジェクトに格納する
func (a *Article) FetchArticles() {

	// Webサイトへのリクエスト
	res, err := http.Get(hokurikuCarURL)
	if err != nil {
		log.Fatal("Http Request Error: %s", err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatal("Response from website is not 200: %d %s", res.StatusCode, res.Status)
	}

	// HTMLドキュメントの取得
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal("Load HTML Document Error: %s", err)
	}
	doc.Find(articleSelectorPath).Each(func(i int, s *goquery.Selection) {
		// 取得できた記事１つ１つに対する処理
		title := s.Text()
		url, _ := s.Attr("href")

		if i == 3 {
			a.Title = title
			a.URL = url
		}
	})
	log.Println("%+v", a)
}
