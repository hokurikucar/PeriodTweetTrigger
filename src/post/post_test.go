package post

import (
	"reflect"
	"testing"
)

func TestNewPostObject(t *testing.T) {
	tests := []struct {
		name string
		want *Post
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPostObject(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPostObject() = %v, want %v", got, tt.want)
			}
		})
	}
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
		// TODO: Add test cases.
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
	tests := []struct {
		name string
		want int
	}{
		// TODO: Add test cases.
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
	tests := []struct {
		name string
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getPostIndexNumber(); got != tt.want {
				t.Errorf("getPostIndexNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
