package routes

import (
	"net/http"

	"github.com/emmanuelbenson/gpi-validate-v2/apis/v2/educert"
	"github.com/emmanuelbenson/gpi-validate-v2/apis/v2/prevemployment"
	"github.com/go-chi/chi"
)

// InitRouter initializes all api routes
func InitRouter() http.Handler {
	r := chi.NewRouter()

	r.Route("/api/v2", func(r chi.Router) {
		r.Mount("/educational-certificate", educert.Router())
		r.Mount("/previous-employment", prevemployment.Router())
	})

	return r
}
