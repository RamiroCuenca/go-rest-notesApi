package main

import (
	"github.com/RamiroCuenca/go-rest-notesApi/common/logger"
)

func main() {
	// Init Zap logger so that we can use it all over the app
	logger.InitZapLogger()

	// Init postgres database
	// connection.NewPostgresClient()

	// Get routes
	mux := Routes()

	// Init server
	sv := NewServer(mux)

	// Run server
	logger.Log().Info("Server running over port :8000 ...")
	sv.Run()
}
