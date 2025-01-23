package utils

import (
	"log"
	"os"
)

// Logger is the centralized logger for the application.
var Logger *log.Logger

// InitializeLogger sets up the logger with appropriate settings.
func InitializeLogger() {
	logFile, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	Logger = log.New(logFile, "LOG: ", log.Ldate|log.Ltime|log.Lshortfile)
	Logger.Println("Logger initialized")
}
