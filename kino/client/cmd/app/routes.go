package main

import (
	"net/http"

	"github.com/bmizerany/pat"
)

func (app *App) Routes() http.Handler {
	mux := pat.New()

	mux.Get("/film/id=:id", isAuthorized(app.getFilm))
	mux.Get("/film/add", isAuthorized(app.setFilm))
	mux.Get("/film/all", isAuthorized(app.getFilmList))

	return LogRequest(SecureHeaders(mux))
}
