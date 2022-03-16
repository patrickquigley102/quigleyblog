package quigleyblog

import (
	"embed"
	"html/template"
	"io"
)

var (
	//go:embed "templates/*"
	// PostTemplates are the HTML templates
	PostTemplates embed.FS
)

// PostRenderer renders HTML from a Post
type PostRenderer struct {
	Template *template.Template
}

// NewPostRenderer returns a pointer to a PostRenderer
func NewPostRenderer() (*PostRenderer, error) {
	template, err := template.ParseFS(PostTemplates, "templates/*.tmpl")
	if err != nil {
		return nil, err
	}

	return &PostRenderer{Template: template}, nil
}

// Render converts Post to html and writes it.
func (r *PostRenderer) Render(w io.Writer, p Post) error {
	err := r.Template.Execute(w, p)
	if err != nil {
		return err
	}
	return nil
}
