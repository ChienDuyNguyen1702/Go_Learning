package handlers

import (
	"github.com/ChienDuyNguyen1702/Go_Learning/pkg/config"
	"net/http"

	"github.com/ChienDuyNguyen1702/Go_Learning/pkg/render"
)

var Repo *Repository

// Repo name of the repository used by the handler
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new Repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandler sets the repository for the handler
func NewHandler(r *Repository) {
	Repo = r
}
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
