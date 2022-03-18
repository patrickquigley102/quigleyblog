package quigleyblog

import (
	"bytes"
	"errors"
	"html/template"
	"io"
	"reflect"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
)

func TestNewPostRenderer(t *testing.T) {
	tests := []struct {
		name    string
		want    *PostRenderer
		wantErr bool
	}{
		{
			"create PostRenderer",
			&PostRenderer{Template: templ},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPostRenderer()
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPostRenderer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPostRenderer() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestPostRenderer_RenderIndex(t *testing.T) {
	type args struct {
		w     testWriter
		posts []Post
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"simple posts",
			args{
				w:     &bytes.Buffer{},
				posts: []Post{{Title: "1"}, {Title: "2"}},
			},
			false,
		},
		{
			"title needs santising",
			args{
				w:     &bytes.Buffer{},
				posts: []Post{{Title: "Has Space"}},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &PostRenderer{
				Template: templ,
			}
			err := r.RenderIndex(tt.args.w, tt.args.posts)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostRenderer.RenderIndex() e=%v, wantE %v", err, tt.wantErr)
				return
			}
			approvals.VerifyString(t, tt.args.w.String())
		})
	}
}

func TestPostRenderer_RenderPost(t *testing.T) {
	type args struct {
		w testWriter
		p Post
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"single post",
			args{
				w: &bytes.Buffer{},
				p: Post{
					Title:       "1",
					Description: "A",
					Tags:        []string{"a", "b"},
					Body:        "Body\n",
				},
			},
			false,
		},
		{
			"writer errors",
			args{
				w: &errWriter{},
				p: Post{},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &PostRenderer{
				Template: templ,
			}
			err := r.RenderPost(tt.args.w, tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostRenderer.RenderPost() e=%v, wantE %v", err, tt.wantErr)
				return
			}
			approvals.VerifyString(t, tt.args.w.String())
		})
	}
}

func BenchmarkRenderPost(b *testing.B) {
	var (
		aPost = Post{
			Title:       "1",
			Description: "A",
			Tags:        []string{"a", "b"},
			Body:        "Body\n",
		}
		r = &PostRenderer{
			Template: templ,
		}
	)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r.RenderPost(io.Discard, aPost)
	}
}

var templ, err = template.ParseFS(
	PostTemplates,
	"templates/*.tmpl",
)

type testWriter interface {
	io.Writer
	String() string
}

type errWriter struct{}

func (w errWriter) Write(p []byte) (int, error) {
	return 0, errors.New("")
}

func (w *errWriter) String() string {
	return ""
}
