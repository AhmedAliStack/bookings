package main

import (
	"log"
	"net/http"
	"time"

	"github.com/AhmedAliStack/bookings/pkg/config"
	"github.com/AhmedAliStack/bookings/pkg/handlers"
	"github.com/alexedwards/scs/v2"
)

const portNumber string = ":8080"
var appConfig config.AppConfig
var session *scs.SessionManager

func main() {
	tc := handlers.CreateTemplateCach()
	appConfig.TemplateCach = tc
	appConfig.UseCach = false
	appConfig.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = appConfig.InProduction

	appConfig.AppSession = session
	
	handlers.SetConfig(&appConfig)

	repo := handlers.NewRepo(&appConfig)
	handlers.NewHandlers(repo)

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	// http.ListenAndServe(portNumber, nil)

	srv := http.Server{
		Addr: portNumber,
		Handler: routes(&appConfig),
	}

	err := srv.ListenAndServe()
	if(err != nil) {
		log.Fatal("Error :",err)
	}

}
