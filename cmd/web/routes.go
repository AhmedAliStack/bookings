package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/AhmedAliStack/bookings/pkg/config"
	"github.com/AhmedAliStack/bookings/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(WriteToConsole)
	mux.Use(NoSurf)
	mux.Use(UseSession)

	mux.Get("/",handlers.Repo.Home)
	mux.Get("/about",handlers.Repo.About)
	mux.Get("/generals-quarters",handlers.Repo.Generals)
	mux.Get("/majors-suite",handlers.Repo.Majors)
	mux.Get("/search-availability",handlers.Repo.Availability)
	mux.Post("/search-availability",handlers.Repo.PostAvailability)
	mux.Get("/make-reservation",handlers.Repo.Reservation)
	mux.Get("/contact",handlers.Repo.Contact)
	mux.Post("/search-availability-json",handlers.Repo.AvailabilityJson)
	
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*",http.StripPrefix("/static",fileServer))
	
	return mux
}