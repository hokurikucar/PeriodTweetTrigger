package tweet

import (
	"errors"
	"net/url"
	"os"
	"strings"

	"github.com/ChimeraCoder/anaconda"
)

// TwitterAPICredentialGenerator PostTweetメソッドをWrapするためのinterface
type TwitterAPICredentialGenerator interface {
	PostTweet(string, url.Values) (anaconda.Tweet, error)
}

// tw TwitterAPIのCredential情報を保有しているオブジェクトが代入されている変数
var tw TwitterAPICredentialGenerator = anaconda.NewTwitterApiWithCredentials(
	os.Getenv("TWITTER_ACCESS_TOKEN"),
	os.Getenv("TWITTER_ACCESS_TOKEN_SECRET"),
	os.Getenv("TWITTER_API_KEY"),
	os.Getenv("TWITTER_SECRET_KEY"),
)

// Tweet APIにて連携済みのtwitterアカウントにツイートを送信する
func Tweet(text string, url string, tags []string) error {
	// 記事にタグをつけない可能性があるので、タグの空判定は行わない
	if text == "" || url == "" {
		return errors.New("Invalid parameter received")
	}
	tagContent := strings.Join(tags, " ")
	tweetContent := text + "\n" + url + "\n" + tagContent
	_, err := tw.PostTweet(tweetContent, nil)
	if err != nil {
		return err
	}
	return nil
}
