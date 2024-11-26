package routes

import (
	"simple_api/handlers"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r chi.Router) {
	mux := chi.NewRouter()
	mux.Route("/v1", func(r chi.Router) {
		r.Get("/health", handlers.HealthCheck)
		r.Route("/users", func(r chi.Router) {
			r.Get("/", handlers.GetUsers)
			r.Post("/", handlers.CreateUser)
			r.Get("/{id}", handlers.GetUser)
		})
	})

	r.Mount("/api", mux)
}
