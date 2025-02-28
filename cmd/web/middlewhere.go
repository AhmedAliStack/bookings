package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Page Loading")
		next.ServeHTTP(w,r)
	})
}

func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Secure: appConfig.InProduction,
		Path: "/",
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

func UseSession(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
