package handlers

import (
	"net/http"

	"github.com/ggratte/modern-web-apps-with-go-udemy/pkg/config"
	"github.com/ggratte/modern-web-apps-with-go-udemy/pkg/models"
	"github.com/ggratte/modern-web-apps-with-go-udemy/pkg/render"
)

const remoteIPKey = "remote_ip"

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{a}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), remoteIPKey, remoteIP)
	render.RenderTemplate(w, "home.page.tmpl.html", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"

	remoteIP := m.App.Session.GetString(r.Context(), remoteIPKey)
	stringMap[remoteIPKey] = remoteIP

	// send the data
	render.RenderTemplate(w, "about.page.tmpl.html", &models.TemplateData{StringMap: stringMap})
}
