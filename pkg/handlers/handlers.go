package handlers

import (
	"net/http"

	"github.com/olive-okamoto/go-startup/pkg/config"
	"github.com/olive-okamoto/go-startup/pkg/models"
	"github.com/olive-okamoto/go-startup/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// Creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// Sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplete(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "hello again"

	render.RenderTemplete(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
