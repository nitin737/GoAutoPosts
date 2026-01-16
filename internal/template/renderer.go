package template

import (
	"bytes"
	"embed"
	"fmt"
	"text/template"

	"github.com/nitin737/GoAutoPosts/internal/model"
)

//go:embed *.tmpl
var templates embed.FS

// Renderer handles template rendering
type Renderer struct {
	templates *template.Template
}

// NewRenderer creates a new template renderer
func NewRenderer() (*Renderer, error) {
	tmpl, err := template.ParseFS(templates, "*.tmpl")
	if err != nil {
		return nil, fmt.Errorf("failed to parse templates: %w", err)
	}

	return &Renderer{
		templates: tmpl,
	}, nil
}

// RenderCaption renders the Instagram caption for a library
func (r *Renderer) RenderCaption(lib *model.Library, hashtags []string) (string, error) {
	var buf bytes.Buffer

	data := map[string]interface{}{
		"Library":  lib,
		"Hashtags": hashtags,
	}

	if err := r.templates.ExecuteTemplate(&buf, "caption.tmpl", data); err != nil {
		return "", fmt.Errorf("failed to render caption: %w", err)
	}

	return buf.String(), nil
}
