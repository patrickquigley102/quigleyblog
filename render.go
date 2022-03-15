package quigleyblog

import (
	"embed"
	"html/template"
	"io"
)

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

// Render converts Post to html and writes it.
func Render(w io.Writer, p Post) error {
	template, err := template.ParseFS(postTemplates, "templates/*.tmpl")
	if err != nil {
		return err
	}

	err = template.Execute(w, p)
	if err != nil {
		return err
	}
	return nil
}
