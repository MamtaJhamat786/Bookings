package handlers

import (
	"github.com/MamtaJhamat786/Bookings/pkg/config"
	"github.com/MamtaJhamat786/Bookings/pkg/models"
	"github.com/MamtaJhamat786/Bookings/pkg/render"
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

// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remote_ip := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remote_ip)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the handler for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"

	remote_ip := m.App.Session.GetString(r.Context(), "remote_ip")

	
	stringMap["remoteIP"] = remote_ip

	// send data to the template
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
