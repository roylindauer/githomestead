package main

import (
	"gitapi/internal/gitserver"
	"gitapi/internal/transport"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	svc := gitserver.NewService(os.Getenv("GIT_SERVER_ENDPOINT"))

	server := transport.NewServer(*svc)
	if err := server.Serve(); err != nil {
		log.Println(err)
	}
}
