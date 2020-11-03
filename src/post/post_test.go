package post

import (
	"errors"
	"reflect"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

type RandomNumberGeneratorMock struct {
	MockNumber int
}

// Intn randパッケージのIntnメソッドのモック
func (rng *RandomNumberGeneratorMock) Intn(n int) int {
	result := rng.MockNumber
	return result
}

// 記事の取得に関する動作のInterfaceをモック
type FetcherMock struct {
}

// 記事の取得に関する動作をモックしたメソッド群
func (f *FetcherMock) choosePostURLRandomly() (string, error) {
	return "dummyURL", nil
}
func (f *FetcherMock) getTitleAndTags(url string) (string, []string, error) {
	if url == "error" {
		return "", nil, errors.New("mock error occured")
	}
	var dummyTags = []string{"hoge", "fuga"}
	return "dummyTitle", dummyTags, nil
}
func TestPost_FetchPosts(t *testing.T) {
	fm := &FetchWorker{&FetcherMock{}}
	type fields struct {
		Title string
		URL   string
		Tags  []string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "記事のURLが指定されなかった場合にランダムにURLを取得すること",
			fields: fields{
				Title: "",
				URL:   "",
				Tags:  nil,
			},
			wantErr: false,
		},
		{
			name: "予め記事のURLが指定されている場合には、そのURLを基にタイトルとタグを取得すること",
			fields: fields{
				Title: "",
				URL:   "dummyURL",
				Tags:  nil,
			},
			wantErr: false,
		},
		{
			name: "タイトルとタグの取得に失敗したときにエラーを返すこと",
			fields: fields{
				Title: "",
				URL:   "error",
				Tags:  nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Post{
				Title: tt.fields.Title,
				URL:   tt.fields.URL,
			}
			if err := fm.FetchPosts(p); (err != nil) != tt.wantErr {
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

func TestPost_choosePostURLRandomly(t *testing.T) {
	rng = &RandomNumberGeneratorMock{MockNumber: 1}
	type fields struct {
		Title string
		URL   string
		Tags  []string
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Post{
				Title: tt.fields.Title,
				URL:   tt.fields.URL,
				Tags:  tt.fields.Tags,
			}
			got, err := p.choosePostURLRandomly()
			if (err != nil) != tt.wantErr {
				t.Errorf("Post.choosePostURLRandomly() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Post.choosePostURLRandomly() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPost_getTitleAndTags(t *testing.T) {
	type fields struct {
		Title string
		URL   string
		Tags  []string
	}
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		want1   []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Post{
				Title: tt.fields.Title,
				URL:   tt.fields.URL,
				Tags:  tt.fields.Tags,
			}
			got, got1, err := p.getTitleAndTags(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("Post.getTitleAndTags() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Post.getTitleAndTags() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Post.getTitleAndTags() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_fetchTags(t *testing.T) {
	type args struct {
		d *goquery.Document
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fetchTags(tt.args.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fetchTags() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_execQuery(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		want    *goquery.Document
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := execQuery(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("execQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("execQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}
