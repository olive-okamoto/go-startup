package main

import (
	"fmt"
	"net/http"
	"text/template"
)

const portNumber = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplete(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplete(w, "about.page.tmpl")
}

func renderTemplete(w http.ResponseWriter, tmpl string) {
	parsedTemplete, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplete.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template", err)
	}
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	_ = http.ListenAndServe(portNumber, nil)
}
