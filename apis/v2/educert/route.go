package educert

import (
	"net/http"

	"github.com/go-chi/chi/middleware"

	"github.com/go-chi/chi"
)

// Router is Educational Certificate API router
func Router() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/", func(r chi.Router) {
		r.Post("/", All)
		r.Post("/new", New)
		r.Get("/show", Show)
	})

	return r
}
