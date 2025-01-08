package main

import (
	"gitapi/internal/gitserver"
	"gitapi/internal/transport"
	"log"
)

func main() {
	svc := gitserver.NewService()

	server := transport.NewServer(*svc)
	if err := server.Serve(); err != nil {
		log.Println(err)
	}
}
