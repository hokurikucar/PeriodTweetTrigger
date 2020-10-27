package tweet

import (
	"net/url"
	"os"

	"github.com/ChimeraCoder/anaconda"
)

// Credential ツイートに必要な情報を持つオブジェクト
type Credential struct {
	credit *anaconda.TwitterApi
}

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
func Tweet(text string, url string) error {
	tweetContent := text + "\n" + url
	_, err := tw.PostTweet(tweetContent, nil)
	if err != nil {
		return err
	}
	return nil
}
