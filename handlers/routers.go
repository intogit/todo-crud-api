package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func CreateRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Get("/healthcheck", HandleHeathCheck)
			r.Post("/todo/create", HandleCreateTodo)
			r.Get("/todo/getAll", HandleGetAllTodo)
			r.Get("/todo/GetById/{id}", HandleGetTodoById)
		})
		r.Route("/v2", func(r chi.Router) {
			r.Get("/healthcheck", HandleHeathCheck)
		})
	})

	return r
}
