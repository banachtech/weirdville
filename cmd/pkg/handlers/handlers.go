package handlers

import (
	"math/rand"
	"net/http"

	"banachtech.github.com/weirdville/cmd/pkg/config"
	"banachtech.github.com/weirdville/cmd/pkg/models"
	"banachtech.github.com/weirdville/cmd/pkg/render"
)

// Repository pattern
// We create a repository of handlers
// This allows handlers to access app config

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{a}
}

func NewHandlers(repo *Repository) {
	Repo = repo
}

// Home is the home page handler
func (repo *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	repo.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (repo *Repository) About(w http.ResponseWriter, r *http.Request) {
	floatMap := make(map[string]float32)
	stringMap := make(map[string]string)

	// some random numbers
	floatMap["uniform"] = rand.Float32()
	floatMap["normal"] = float32(rand.NormFloat64())

	// retrieve remote IP address from session manager
	remoteIP := repo.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		FloatMap: floatMap,
		StrMap:   stringMap,
	})
}
