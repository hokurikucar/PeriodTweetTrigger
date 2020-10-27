package queryFetcher

import (
	"math/rand"
	"net/http"
	"strconv"
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
	pagenationIndex := getPagenationNumber()
	articleIndex := getArticleIndexNumber()
	// Webサイトへのリクエスト
	res, err := http.Get(hokurikuCarURL + "/page/" + strconv.Itoa(pagenationIndex))
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
		return err
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

// getPagenationNumber どのページの記事を取得するかを決定する
func getPagenationNumber() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(4)
}

// getArticleIndexNumber どのインデックス番号の記事を取得するかを決定する
func getArticleIndexNumber() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(9)
}
