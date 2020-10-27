package tweet

import (
	"os"

	"github.com/ChimeraCoder/anaconda"
)

// Credential ツイートに必要な情報を持つオブジェクト
type Credential struct {
	credit *anaconda.TwitterApi
}

// GetTwitterCredential 認証済みTwitterAPIを持つオブジェクトを返却する
func GetTwitterCredential() *Credential {
	var c Credential
	c.credit = anaconda.NewTwitterApiWithCredentials(
		os.Getenv("TWITTER_ACCESS_TOKEN"),
		os.Getenv("TWITTER_ACCESS_TOKEN_SECRET"),
		os.Getenv("TWITTER_API_KEY"),
		os.Getenv("TWITTER_SECRET_KEY"),
	)
	return &c
}

// Tweet APIにて連携済みのtwitterアカウントにツイートを送信する
func (c *Credential) Tweet(text string, url string) error {
	tweetContent := text + "\n" + url
	_, err := c.credit.PostTweet(tweetContent, nil)
	if err != nil {
		return err
	}
	return nil
}
