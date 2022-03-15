package tests

import (
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"

	"github.com/patrickquigley102/quigleyblog"
)

func TestNewPostsFromFS(t *testing.T) {
	type args struct {
		fileSystem fs.FS
	}
	tests := []struct {
		name    string
		args    args
		want    []quigleyblog.Post
		wantErr bool
	}{
		{
			"multiple files",
			args{
				fstest.MapFS{
					"1.md": {
						Data: []byte(
							"Title: 1\nDescription: A\nTags: a,b\n---\nBody",
						),
					},
					"2.md": {
						Data: []byte(
							"Title: 2\nDescription: B\nTags: c,d\n---\nBody\nmorebody",
						),
					},
				},
			},
			[]quigleyblog.Post{
				{
					Title:       "1",
					Description: "A",
					Tags:        []string{"a", "b"},
					Body:        "Body\n",
				},
				{
					Title:       "2",
					Description: "B",
					Tags:        []string{"c", "d"},
					Body:        "Body\nmorebody\n",
				},
			},
			false,
		},
		{
			"error opening fs",
			args{stubFailingFS{}},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := quigleyblog.NewPostsFromFS(tt.args.fileSystem)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPostsFromFS() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPostsFromFS() = %v, want %v", got, tt.want)
			}
		})
	}

}

type stubFailingFS struct{}

func (s stubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("")
}
