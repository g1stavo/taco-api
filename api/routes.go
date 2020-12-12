package api

import (
	"net/http"

	"github.com/go-chi/chi"
)

// NewRouter creates new HTTP router
func NewRouter() http.Handler {
	r := chi.NewRouter()
	r.Use(Logger)
	r.Get("/taco", getTaco)
	return r
}
