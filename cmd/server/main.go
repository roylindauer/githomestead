package main

import (
	"gitapi/internal/gitserver"
	"gitapi/internal/transport"
	"github.com/go-yaml/yaml"
	"log"
	"os"
)

type Config struct {
	Port     string `yaml:"port"`
	Endpoint string `yaml:"endpoint"`
}

func main() {
	f, err := os.Open("config.yml")
	if err != nil {
		log.Println(err)
		panic(err)
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Println(err)
			panic(err)
		}
	}(f)

	var cfg Config

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	svc := gitserver.NewService(cfg.Endpoint)

	server := transport.NewServer(*svc)
	if err := server.Serve(); err != nil {
		log.Println(err)
	}
}
