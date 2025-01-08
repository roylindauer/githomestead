package main

import (
	"fmt"
	"gitapi/internal/gitserver"
	"gitapi/internal/transport"
)

func main() {
	svc := gitserver.NewService()

	server := transport.NewServer(*svc)
	if err := server.Serve(); err != nil {
		fmt.Println(err)
	}
}
