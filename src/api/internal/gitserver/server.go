package gitserver

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Service struct {
	Url      string
	RepoPath string
}

type Repo struct {
	Url         string
	Description string
	Name        string
}

func NewService() *Service {
	return &Service{
		Url:      "http://localhost:9184",
		RepoPath: "../../repos",
	}
}

func (s Service) GetAllRepos() []Repo {
	// so i basically just need to read from /repos and return a list of .git projects
	root := s.RepoPath
	var repos []Repo

	directories, err := os.ReadDir(root)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range directories {
		fmt.Println(filepath.Join(root, file.Name()))
		if file.IsDir() && strings.HasSuffix(file.Name(), ".git") {
			strings.HasSuffix(file.Name(), ".git")

			url := fmt.Sprintf("%s/%s", s.Url, file.Name())
			description, err := os.ReadFile(filepath.Join(root, file.Name(), "description"))

			if err != nil {
				log.Fatal(err)
			}

			repo := Repo{
				Name:        file.Name(),
				Url:         url,
				Description: strings.TrimSpace(string(description)),
			}

			repos = append(repos, repo)
		}
	}

	return repos
}
