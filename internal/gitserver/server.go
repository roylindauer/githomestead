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
	var repos []Repo

	directories, err := os.ReadDir(s.RepoPath)

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range directories {
		fmt.Println(filepath.Join(s.RepoPath, file.Name()))
		if file.IsDir() && strings.HasSuffix(file.Name(), ".git") {
			strings.HasSuffix(file.Name(), ".git")

			description, err := os.ReadFile(filepath.Join(s.RepoPath, file.Name(), "description"))

			if err != nil {
				log.Fatal(err)
			}

			repo := Repo{
				Name:        file.Name(),
				Url:         fmt.Sprintf("%s/%s", s.Url, file.Name()),
				Description: strings.TrimSpace(string(description)),
			}

			repos = append(repos, repo)
		}
	}

	return repos
}

func (s Service) Create() {

}

func (s Service) Update() {

}

func (s Service) Destroy() {

}

func (s Service) Get() {

}
