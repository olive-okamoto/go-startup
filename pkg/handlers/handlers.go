package handlers

import (
	"net/http"

	"github.com/olive-okamoto/go-startup/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplete(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplete(w, "about.page.tmpl")
}
