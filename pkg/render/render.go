package render

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

var function = template.FuncMap{}

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	_, err := RenderTemplateTest(w)
	if err != nil {
		fmt.Println("Error getting template catche", err)
	}

	parsedTemp, _ := template.ParseFiles("./templates/" + tmpl)
	err = parsedTemp.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template", err)
	}
}

func RenderTemplateTest(w http.ResponseWriter) (map[string]*template.Template, error) {
	myCatche := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")

	if err != nil {
		return myCatche, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		fmt.Println("Page is currently", page)
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
