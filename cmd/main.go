package main

import (
	"log"

	"github.com/rizhyi/6-sprint-final/internal/server"
)

func main() {

	logger := log.New(log.Writer(), "LOG: ", log.Ldate|log.Ltime|log.Lshortfile)

	server := server.CreateServer(logger)

	logger.Println("Запуск сервера на порту 8080")
	if err := server.Server.ListenAndServe(); err != nil {
		logger.Fatalf("Ошибка при запуске сервера: %v", err)
	}
}
