package handlers

import (
	"github.com/tnaucoin/go-web-app/pkg/config"
	"github.com/tnaucoin/go-web-app/pkg/models"
	"github.com/tnaucoin/go-web-app/pkg/render"
	"net/http"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.TemplateRenderer(w, "home.page.html", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Test_Data"
	render.TemplateRenderer(w, "about.page.html", &models.TemplateData{StringMap: stringMap})
}
