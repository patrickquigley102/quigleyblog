package quigleyblog

import (
	"embed"
	"html/template"
	"io"
)

const (
	templatesRegex = "templates/*.tmpl"
	postTemplate   = "post.tmpl"
	indexTemplate  = "index.tmpl"
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
	template, err := template.ParseFS(PostTemplates, templatesRegex)
	if err != nil {
		return nil, err
	}

	return &PostRenderer{Template: template}, nil
}

// RenderPost converts Post to html and writes it.
func (r *PostRenderer) RenderPost(w io.Writer, p Post) error {
	err := r.Template.ExecuteTemplate(w, postTemplate, p)
	if err != nil {
		return err
	}
	return nil
}

// RenderIndex renders the html of an index of Posts
func (r *PostRenderer) RenderIndex(w io.Writer, posts []Post) error {
	err := r.Template.ExecuteTemplate(w, indexTemplate, posts)
	if err != nil {
		return err
	}
	return nil
}
