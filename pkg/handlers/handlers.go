package handlers

import (
	"encoding/json"
	"fmt"
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
	renderTemplate(w, "home.page.tmpl", &models.TemplateData{}, r)
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	remoteIp := appConfig.AppSession.GetString(r.Context(),"remote_ip")
	stringMap := map[string]string{
		"test" : "Hello from Go",
		"remote_ip" : remoteIp,
	}
	renderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	}, r)
}

func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "make-reservation.page.tmpl", &models.TemplateData{}, r)
}

func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "generals.page.tmpl", &models.TemplateData{}, r)
}

func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "majors.page.tmpl", &models.TemplateData{}, r)
}

func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "search-availability.page.tmpl", &models.TemplateData{}, r)
}

type jsonResponse struct {
	Ok bool `json:"ok"`
	Message string `json:"message"`
}

func (m *Repository) AvailabilityJson(w http.ResponseWriter, r *http.Request) {
	response := jsonResponse{
		Ok: true,
		Message: "Available!",
	}
	
	out, err := json.MarshalIndent(response, "", "    ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.FormValue("start")
	end := r.FormValue("end")
	w.Write([]byte(fmt.Sprintf("start date is %s and end date is %s", start, end)))
}

func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "contact.page.tmpl", &models.TemplateData{}, r)
}
