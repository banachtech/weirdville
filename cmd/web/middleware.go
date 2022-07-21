package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// create CSRF token
func NoSurf(h http.Handler) http.Handler {
	csrfHandler := nosurf.New(h)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// load Session
func LoadSession(h http.Handler) http.Handler {
	return session.LoadAndSave(h)
}
