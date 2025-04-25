package server

import (
	"log"
	"net/http"
	"time"

	"github.com/rizhyi/6-sprint-final/internal/handlers"
)

// Server structure
type CustomServer struct {
	Logger *log.Logger
	Server *http.Server
}

func CreateServer(logger *log.Logger) *CustomServer {
	// Creating router
	router := http.NewServeMux()

	// Reg handlers
	router.HandleFunc("/", handlers.HandlerMain)
	router.HandleFunc("/upload", handlers.HandlerUpload)

	// Creating server
	httpServer := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	// Returning server
	return &CustomServer{
		Logger: logger,
		Server: httpServer,
	}
}
