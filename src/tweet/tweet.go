package tweet

import (
	"os"

	"github.com/ChimeraCoder/anaconda"
)

// Tweet APIにて連携済みのtwitterアカウントにツイートを送信する
func Tweet(text string, url string) error {
	tweetContent := text + "\n" + url
	api := getCredential()
	_, err := api.PostTweet(tweetContent, nil)
	if err != nil {
		return err
	}
	return nil
}

// getCredential twitterアカウントのAPIを取得する
func getCredential() *anaconda.TwitterApi {
	return anaconda.NewTwitterApiWithCredentials(
		os.Getenv("TWITTER_ACCESS_TOKEN"),
		os.Getenv("TWITTER_ACCESS_TOKEN_SECRET"),
		os.Getenv("TWITTER_API_KEY"),
		os.Getenv("TWITTER_SECRET_KEY"),
	)
}
