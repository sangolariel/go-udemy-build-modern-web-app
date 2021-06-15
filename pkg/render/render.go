package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/sangolariel/go-udemy-build-modern-web-app/pkg/config"
)

var function = template.FuncMap{}

var app *config.AppConfig

func NewTemplates(config *config.AppConfig) {
	app = config
}

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	var tc map[string]*template.Template
	if app.UseCatche {
		tc = app.TemplateCatche
	} else {
		tc, _ = CreateTemplateCatche()
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get a template from template catche")
	}
	buf := new(bytes.Buffer)

	_ = t.Execute(buf, nil)

	_, _ = buf.WriteTo(w)

}

func CreateTemplateCatche() (map[string]*template.Template, error) {
	myCatche := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")

	if err != nil {
		return myCatche, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(function).ParseFiles(page)
		if err != nil {
			return myCatche, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCatche, err
		}

		if len(matches) > 0 {
			ts, err := ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCatche, err
			}
			myCatche[name] = ts
		}
	}
	return myCatche, err
}
