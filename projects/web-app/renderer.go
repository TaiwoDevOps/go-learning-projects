package main

import (
	"html/template"
	"net/http"
	"path"
	"path/filepath"
	"sync"
)

type TemplateRenderer struct {
	cache       map[string]*template.Template
	mutex       sync.RWMutex
	dev         bool
	templateDir string
}

type templateData struct {
	Form            *Form
	IsAuthenticated bool
	Flash           string
	Posts           []Post
	Metadata        Metadata
	Comments        []Comment
	Post            *Post
	NextLink        string
	PreviousLink    string
}

func NewTemplateRender(templateDir string, isDev bool) *TemplateRenderer {
	return &TemplateRenderer{
		cache:       make(map[string]*template.Template),
		dev:         isDev,
		templateDir: templateDir,
	}
}

func (t *TemplateRenderer) Render(w http.ResponseWriter, templateName string, data interface{}) {
	templ, err := t.getTemplate(templateName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = templ.ExecuteTemplate(w, "base.html", data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (t *TemplateRenderer) getTemplate(templateName string) (*template.Template, error) {
	if !t.dev {
		t.mutex.RLock()
		if templ, ok := t.cache[templateName]; ok {
			t.mutex.RUnlock()
			return templ, nil
		}
		t.mutex.RUnlock()
	}

	templ, err := t.parseTemplate(templateName)
	if err != nil {
		return nil, err
	}

	if !t.dev {
		t.mutex.Lock()
		t.cache[templateName] = templ
		t.mutex.Unlock()
	}

	return templ, nil
}

func (t *TemplateRenderer) parseTemplate(templateName string) (*template.Template, error) {
	templatePath := path.Join(t.templateDir, templateName)
	files := []string{templatePath}
	files = t.parseTemplateJoiner("layouts/*.html", files)
	files = t.parseTemplateJoiner("partials/*.html", files)

	templ, err := template.ParseFiles(files...)
	if err != nil {
		return nil, err
	}

	return templ, nil
}

func (t *TemplateRenderer) parseTemplateJoiner(p string, files []string) []string {
	pathResult := path.Join(t.templateDir, p)
	result, err := filepath.Glob(pathResult)
	if err == nil {
		files = append(files, result...)
	}
	return files
}
