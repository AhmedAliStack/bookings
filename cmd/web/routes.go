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

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*",http.StripPrefix("/static",fileServer))
	
	return mux
}