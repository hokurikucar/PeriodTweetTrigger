package queryFetcher

import (
	"log"
	"math/rand"
	"net/http"
	"time"

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
func (a *Article) FetchArticles() error {

	// 取得する記事の番号をランダムに設定
	rand.Seed(time.Now().UnixNano())
	articleIndex := rand.Intn(9)

	// Webサイトへのリクエスト
	res, err := http.Get(hokurikuCarURL)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return err
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

		if i == articleIndex {
			a.Title = title
			a.URL = url
		}
	})
	return nil
}
