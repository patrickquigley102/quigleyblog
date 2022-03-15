package tests

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/patrickquigley102/quigleyblog"
)

func TestRender(t *testing.T) {
	type args struct {
		w bytes.Buffer
		p quigleyblog.Post
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			"single post",
			args{
				w: bytes.Buffer{},
				p: quigleyblog.Post{
					Title:       "1",
					Description: "A",
					Tags:        []string{"a", "b"},
					Body:        "Body\n",
				},
			},
			"<h1>1</h1>\n\n<p>A</p>\n\nTags: <ul><li>a</li><li>b</li></ul>\n",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := quigleyblog.Render(&tt.args.w, tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("Render() error = %v, wantErr %v", err, tt.wantErr)
			}
			got := tt.args.w.String()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Render() = %v, want %v", got, tt.want)
			}
		})
	}
}
