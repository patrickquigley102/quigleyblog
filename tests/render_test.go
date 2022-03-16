package tests

import (
	"bytes"
	"errors"
	"io"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
	"github.com/patrickquigley102/quigleyblog"
)

func TestRender(t *testing.T) {
	type args struct {
		w testWriter
		p quigleyblog.Post
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
				p: quigleyblog.Post{
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
				p: quigleyblog.Post{},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := quigleyblog.Render(tt.args.w, tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("Render() error = %v, wantErr %v", err, tt.wantErr)
			}
			approvals.VerifyString(t, tt.args.w.String())
		})
	}
}

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

func BenchmarkRender(b *testing.B) {
	var (
		aPost = quigleyblog.Post{
			Title:       "1",
			Description: "A",
			Tags:        []string{"a", "b"},
			Body:        "Body\n",
		}
	)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		quigleyblog.Render(io.Discard, aPost)
	}
}
