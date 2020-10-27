package tweet

import (
	"reflect"
	"testing"

	"github.com/ChimeraCoder/anaconda"
)

func TestTweet(t *testing.T) {
	type args struct {
		text string
		url  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Tweet(tt.args.text, tt.args.url); (err != nil) != tt.wantErr {
				t.Errorf("Tweet() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_getCredential(t *testing.T) {
	tests := []struct {
		name string
		want *anaconda.TwitterApi
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCredential(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getCredential() = %v, want %v", got, tt.want)
			}
		})
	}
}
