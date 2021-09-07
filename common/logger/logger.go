package logger

import (
	"log"

	"go.uber.org/zap"
)

var sugar *zap.SugaredLogger

// This function inits the zap logger.
// It may be called from the main package at the start of the application.
func InitZapLogger() error {
	l, err := zap.NewDevelopment()
	if err != nil {
		log.Printf("Coul not init logger. Reason: %v", err)
		return err
	}

	sugar = l.Sugar()

	return nil
}

// Through this function we can use the suggared logger from zap package
//
// We call it Log cause the package it's already logger and we dont want to be redundant.
func Log() *zap.SugaredLogger {
	return sugar
}
