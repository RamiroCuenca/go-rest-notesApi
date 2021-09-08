package main

import (
	"net/http"
	"time"

	"github.com/RamiroCuenca/go-rest-notesApi/common/logger"
	"github.com/go-chi/chi"
)

// We create MyServer struct in order to add the run method to it
type MyServer struct {
	server *http.Server
}

// This function sets up the server configuration and returns the same.
// It receives as a parameter the multiplexer (In this case from chi).
func NewServer(mux *chi.Mux) *MyServer {
	s := &http.Server{
		Addr:           ":8000",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return &MyServer{s}
}

// Run works as a method of MyServer struct an it function is to
// run the server
func (s *MyServer) Run() {
	// By default, the method ListenAndServe() run just for a period
	// of miliseconds, so... envolving it with the method Fatal (Of the logger)
	// we make sure that it will run at least until it get's an error.
	logger.Log().Fatal(s.server.ListenAndServe())
}
