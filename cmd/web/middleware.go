package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

/**
*	Middleware handles common processing for all request
 */

// Adds CSRF protection to all POST request
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// Loads and saves the session oon every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
