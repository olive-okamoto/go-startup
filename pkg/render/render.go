package render

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/olive-okamoto/go-startup/pkg/config"
)

var functions = template.FuncMap{}

var app *config.AppConfig

// Sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

// Render templates using html/template
func RenderTemplete(w http.ResponseWriter, tmpl string) {

	tc := app.TemplateCache

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}
	buf := new(bytes.Buffer)
	// Execute(io.Writer(outputDestination), <data>)
	_ = t.Execute(buf, nil)
	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template", err)
	}
}

// Creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		// New<fileName>.ParseFiles<templateFile>
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}
	return myCache, nil
}
