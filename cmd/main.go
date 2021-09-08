package main

import (
	"github.com/RamiroCuenca/go-rest-notesApi/common/logger"
	"github.com/RamiroCuenca/go-rest-notesApi/database/connection"
)

func main() {
	// Init Zap logger so that we can use it all over the app
	logger.InitZapLogger()

	// Init postgres database
	connection.NewPostgresClient()
}
