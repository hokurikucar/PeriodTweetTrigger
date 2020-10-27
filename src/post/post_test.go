package post

import (
	"testing"
)

type RandomNumberGeneratorMock struct {
	MockNumber int
}

// Intn randパッケージのIntnメソッドのモック
func (rng *RandomNumberGeneratorMock) Intn(n int) int {
	result := rng.MockNumber
	return result
}

func TestPost_FetchPosts(t *testing.T) {
	type fields struct {
		Title string
		URL   string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name:    "記事を取得すること",
			fields:  fields{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Post{
				Title: tt.fields.Title,
				URL:   tt.fields.URL,
			}
			if err := p.FetchPosts(); (err != nil) != tt.wantErr {
				t.Errorf("Post.FetchPosts() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_getPagenationNumber(t *testing.T) {
	// 乱数生成処理のモックを、post.goのプロパティ変数に仕込む
	// これを行うことで、rand.Intnは必ず指定の数字を返すようになる
	rng = &RandomNumberGeneratorMock{MockNumber: 1}
	tests := []struct {
		name string
		want int
	}{
		{
			name: "生成された整数型の乱数に１を足して返却すること",
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getPagenationNumber(); got != tt.want {
				t.Errorf("getPagenationNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getPostIndexNumber(t *testing.T) {
	rng = &RandomNumberGeneratorMock{MockNumber: 1}
	tests := []struct {
		name string
		want int
	}{
		{
			name: "生成された整数型の乱数を返却すること",
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getPostIndexNumber(); got != tt.want {
				t.Errorf("getPostIndexNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
