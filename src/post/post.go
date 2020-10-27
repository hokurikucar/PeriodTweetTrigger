package post

import (
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const hokurikuCarURL = "https://hokurikucar.com/"
const postSelectorPath = "article > div > h2 > a"

// Post 記事の情報を格納するオブジェクト
type Post struct {
	Title string
	URL   string
}

// NewPostObject 取得した記事情報を格納するオブジェクトを生成して返却する
func NewPostObject() *Post {
	return &Post{}
}

// FetchPosts 記事のタイトルとURLを取得してオブジェクトに格納する
func (p *Post) FetchPosts() error {
	pagenationIndex := getPagenationNumber()
	postIndex := getPostIndexNumber()
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
	doc.Find(postSelectorPath).Each(func(i int, s *goquery.Selection) {
		// 取得できた記事１つ１つに対する処理
		title := s.Text()
		url, _ := s.Attr("href")

		if i == postIndex {
			p.Title = title
			p.URL = url
		}
	})
	return nil
}

// getPagenationNumber どのページの記事を取得するかを決定する
func getPagenationNumber() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(4)
}

// getPostIndexNumber どのインデックス番号の記事を取得するかを決定する
func getPostIndexNumber() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(9)
}
