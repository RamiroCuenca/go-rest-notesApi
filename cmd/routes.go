package main

import (
	"github.com/RamiroCuenca/go-rest-notesApi/notes/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func Routes() *chi.Mux {
	// Create a new multiplexer
	r := chi.NewMux()

	// Use logger middleware from chi
	r.Use(middleware.Logger)

	// Handlers
	r.Post("/api/v1/notes/create", handlers.NotesCreate)
	r.Get("/api/v1/notes/readbyid/{id}", nil)
	r.Get("/api/v1/notes/readall", nil)
	r.Put("/api/v1/notes/update", nil)
	r.Delete("/api/v1/notes/delete/{id}", nil)

	return r
}
