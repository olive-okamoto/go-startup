package render

import (
	"fmt"
	"net/http"
	"text/template"
)

// Render templates using html/template
func RenderTemplete(w http.ResponseWriter, tmpl string) {
	parsedTemplete, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplete.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template", err)
	}
}
