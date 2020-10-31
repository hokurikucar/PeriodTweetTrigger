package post

import (
	"errors"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const hokurikuCarURL = "https://hokurikucar.com/"
const postSelectorPath = "article > div > h2 > a"

// RandomNumberGenerator 引数として与えられた数値の範囲で乱数を生成するインタフェース
type RandomNumberGenerator interface {
	Intn(n int) int
}

// Post 記事の情報を格納するオブジェクト
type Post struct {
	Title string
	URL   string
	Tags  []string
}

var rng RandomNumberGenerator = rand.New(rand.NewSource(time.Now().UnixNano()))

// NewPostObject 取得した記事情報を格納するオブジェクトを生成して返却する
func NewPostObject() *Post {
	return &Post{}
}

// FetchPosts 記事のタイトルとURLを取得してオブジェクトに格納する
func (p *Post) FetchPosts() error {
	// 将来的に、管理者アプリから指定されたURLにリクエストを送る予定
	// 管理者アプリが完成するまで、もしくはアプリ側でURLが指定されていなかった場合に
	// ランダムで記事を取得する関数を呼ぶ
	if err := p.choosePostURLRandomly(); err != nil {
		return err
	}
	// URLより、タイトルとタグの取得を行う
	doc, err := execQuery(p.URL)
	if err != nil {
		return err
	}
	p.Title = doc.Find("div.viral__contents > h1").Text()
	p.Tags = fetchTags(doc)
	return nil
}

// choosePostURLRandomly 北陸くるま情報サイトより、ランダムで１つ記事のURLを取得する
// このメソッドは、管理者アプリよりURLの指定が無かった場合にのみ呼び出される
func (p *Post) choosePostURLRandomly() error {
	pagenationIndex := getPagenationNumber()
	postIndex := getPostIndexNumber()
	// Webサイトへのリクエスト
	doc, err := execQuery(hokurikuCarURL + "/page/" + strconv.Itoa(pagenationIndex))
	if err != nil {
		return err
	}
	doc.Find(postSelectorPath).Each(func(i int, s *goquery.Selection) {
		if i == postIndex {
			p.URL, _ = s.Attr("href")
		}
	})
	return nil
}

// getPagenationNumber どのページの記事を取得するかを決定する
func getPagenationNumber() int {
	return rng.Intn(4) + 1 // 0番のページネーションは存在しないため
}

// getPostIndexNumber どのインデックス番号の記事を取得するかを決定する
func getPostIndexNumber() int {
	return rng.Intn(9)
}

// 指定のHTMLドキュメントより、タグ情報を取得
func fetchTags(d *goquery.Document) []string {
	var tags []string
	d.Find("div.viral__contents > ul > li.icon-tag > a").Each(func(i int, s *goquery.Selection) {
		tags = append(tags, "#"+s.Text())
	})
	return tags
}

// execQuery URLを基にスクレイピングを行い、HTMLドキュメントを取得する
func execQuery(url string) (*goquery.Document, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, errors.New("Bad Request")
	}
	// HTMLドキュメントの取得
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}
	return doc, nil
}
