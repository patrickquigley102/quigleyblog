package tests

import (
	"bytes"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
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
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := quigleyblog.Render(&tt.args.w, tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("Render() error = %v, wantErr %v", err, tt.wantErr)
			}
			approvals.VerifyString(t, tt.args.w.String())
		})
	}
}
