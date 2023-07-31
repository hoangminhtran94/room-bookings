package handlers

import (
	"net/http"

	"github.com/hoangminhtran94/room-bookings/pkg/config"
	"github.com/hoangminhtran94/room-bookings/pkg/models"
	"github.com/hoangminhtran94/room-bookings/pkg/render"
)

// TemplateData hold page data

// Repo the repository used by the handlers
var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}
func (m *Repository) Home(res http.ResponseWriter, req *http.Request) {
	remoteId := req.RemoteAddr
	m.App.Session.Put(req.Context(), "remote_ip", remoteId)
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."
	render.RenderTemplate(res, "home.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

// About page handler
func (m *Repository) About(res http.ResponseWriter, req *http.Request) {
	remoteIp := m.App.Session.GetString(req.Context(), "remote_ip")
	stringMap := make(map[string]string)
	stringMap["remote_ip"] = remoteIp
	render.RenderTemplate(res, "about.html", &models.TemplateData{StringMap: stringMap})
}
