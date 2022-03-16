package quigleyblog

import (
	"bytes"
	"html/template"
	"reflect"
	"testing"
)

func TestNewPostRenderer(t *testing.T) {
	tests := []struct {
		name    string
		want    *PostRenderer
		wantErr bool
	}{
		// TODO: Add test cases.
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

func TestPostRenderer_Render(t *testing.T) {
	type fields struct {
		template *template.Template
	}
	type args struct {
		p Post
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantW   string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &PostRenderer{
				template: tt.fields.template,
			}
			w := &bytes.Buffer{}
			if err := r.Render(w, tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("PostRenderer.Render() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("PostRenderer.Render() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}
