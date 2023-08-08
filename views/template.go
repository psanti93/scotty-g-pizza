package views

import (
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
)

type Template struct {
	view *template.Template
}

func Must(t Template, err error) Template {
	if err != nil {
		panic(err)
	}

	return t
}

func ParseFS(fs fs.FS, patterns ...string) (Template, error) {
	tmpl, err := template.ParseFS(fs, patterns...)

	if err != nil {
		return Template{}, fmt.Errorf("parsing: %v", err)
	}

	return Template{view: tmpl}, nil
}

func (t Template) Execute(w http.ResponseWriter, r *http.Request, data interface{}) {
	w.Header().Set("Content-Type", "text/html")
	err := t.view.Execute(w, data)
	if err != nil {
		fmt.Errorf("exeucting: %v", err)
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
	}
}
