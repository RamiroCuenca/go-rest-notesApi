package main

import "github.com/RamiroCuenca/go-rest-notesApi/common/logger"

func main() {
	// Init Zap logger so that we can use it all over the app
	logger.InitZapLogger()
}
