package render

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemp, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemp.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template", err)
	}
}
