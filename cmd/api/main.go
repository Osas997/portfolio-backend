package main

import (
	"log"

	"github.com/Osas997/go-portfolio/internal/config"
	"github.com/Osas997/go-portfolio/internal/server"
)

func main() {
	config.LoadConfig()
	s := server.NewServer()
	if err := s.Start(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
