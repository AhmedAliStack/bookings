package handlers

import (
	"net/http"

	"github.com/AhmedAliStack/bookings/pkg/config"
	"github.com/AhmedAliStack/bookings/pkg/models"
)

var Repo *Repository

type Repository struct {
	appConfig *config.AppConfig
}

func NewRepo(c *config.AppConfig) *Repository {
	return &Repository{
		appConfig: c,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr
	appConfig.AppSession.Put(r.Context(),"remote_ip",remoteIp)
	renderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	remoteIp := appConfig.AppSession.GetString(r.Context(),"remote_ip")
	stringMap := map[string]string{
		"test" : "Hello from Go",
		"remote_ip" : remoteIp,
	}
	renderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
