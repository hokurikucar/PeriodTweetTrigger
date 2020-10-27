package tweet

import (
	"errors"
	"net/url"
	"testing"

	"github.com/ChimeraCoder/anaconda"
)

// TweetAPICredentialGeneratorMock anacondaのPostTweetをMockのためにDIするオブジェクト
type TweetAPICredentialGeneratorMock struct {
	tweet anaconda.Tweet
	error error
}

// anacondaパッケージのPostTweetメソッドのMock
// エラーを意図的に起こすために、第一引数に特定の文字列を受け取った場合にエラーを生成して返却する
func (tcg *TweetAPICredentialGeneratorMock) PostTweet(c string, n url.Values) (anaconda.Tweet, error) {
	var t anaconda.Tweet
	if c == "invalid tweet \n dummyURL" {
		return t, errors.New("dummy tweet error")
	}
	return t, nil
}

func TestTweet(t *testing.T) {
	var at anaconda.Tweet
	tw = &TweetAPICredentialGeneratorMock{tweet: at, error: nil}
	type args struct {
		text string
		url  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "記事の投稿に成功した場合にnilのエラーを返すこと",
			args:    args{text: "dummy tweet", url: ""},
			wantErr: false,
		},
		{
			name:    "記事の投稿に失敗した場合にエラーを返すこと",
			args:    args{text: "invalid tweeet", url: "dummyURL"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Tweet(tt.args.text, tt.args.url); (err != nil) != tt.wantErr {
				t.Errorf("Tweet() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
