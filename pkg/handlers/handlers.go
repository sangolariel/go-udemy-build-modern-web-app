package handlers

import (
	"net/http"

	"github.com/sangolariel/go-udemy-build-modern-web-app/pkg/config"
	"github.com/sangolariel/go-udemy-build-modern-web-app/pkg/models"
	"github.com/sangolariel/go-udemy-build-modern-web-app/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

// NewRepository create a new repository
func NewRepository(config *config.AppConfig) *Repository {
	return &Repository{
		App: config,
	}
}

//Handler Set a repository for the handler
func NewHandler(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hope to see you again!"
	//send data to the template
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
