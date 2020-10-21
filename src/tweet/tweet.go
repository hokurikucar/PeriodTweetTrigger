package tweet

import (
	"log"
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/joho/godotenv"
)

// Tweet APIにて連携済みのtwitterアカウントにツイートを送信する
func Tweet(text string, url string) error {
	// 環境変数の読み込み
	// TODO build後のディレクトリ構成を調べてenvファイルを読み込む
	if err := godotenv.Load(); err != nil {
		log.Fatal("Load Environment Error: %v", err)
	}

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
	return anaconda.NewTwitterApi(os.Getenv("TWITTER_ACCESS_TOKEN"), os.Getenv("TWITTER_ACCESS_TOKEN_SECRET"))
}
