package main

import (
	"log"

	"github.com/ayushpal11/wav-flac-converter/internal/server"
)

func main() {
	srv := server.NewServer()
	if err := srv.Start(":8080"); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

// File for entry point of the application
