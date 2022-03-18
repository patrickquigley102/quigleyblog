package quigleyblog

import "testing"

func TestPost_SanitizedTitle(t *testing.T) {
	type fields struct {
		Title string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"whitespace",
			fields{
				Title: "a title",
			},
			"a-title",
		},
		{
			"upper case",
			fields{
				Title: "UpperCase",
			},
			"uppercase",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Post{
				Title: tt.fields.Title,
			}
			if got := p.SanitizedTitle(); got != tt.want {
				t.Errorf("Post.SanitizedTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}
