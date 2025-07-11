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
	if c == "invalid tweet\ndummyURL\nhoge fuga" {
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
		tags []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "記事の投稿に成功した場合にnilのエラーを返すこと",
			args:    args{text: "dummy tweet", url: "dummyURL", tags: []string{"hoge", "fuga"}},
			wantErr: false,
		},
		{
			name:    "記事の投稿に失敗した場合にエラーを返すこと",
			args:    args{text: "invalid tweet", url: "dummyURL", tags: []string{"hoge", "fuga"}},
			wantErr: true,
		},
		{
			name:    "記事のタイトルが空文字列だった場合にエラーを返すこと",
			args:    args{text: "", url: "dummyURL", tags: []string{"hoge", "fuga"}},
			wantErr: true,
		},
		{
			name:    "URLが空文字列だった場合にエラーを返すこと",
			args:    args{text: "dummyTweet", url: "", tags: []string{"hoge", "fuga"}},
			wantErr: true,
		},
		{
			name:    "タグが無い場合でもエラーは返さないこと",
			args:    args{text: "dummyTweet", url: "dummyURL", tags: []string{}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Tweet(tt.args.text, tt.args.url, tt.args.tags); (err != nil) != tt.wantErr {
				t.Errorf("Tweet() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
