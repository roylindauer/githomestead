package gitserver

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
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
		Url:      "git://localhost:9418",
		RepoPath: "repos",
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
				Name:        TrimSuffix(file.Name(), ".git"),
				Url:         fmt.Sprintf("%s/%s", s.Url, file.Name()),
				Description: strings.TrimSpace(string(description)),
			}

			repos = append(repos, repo)
		}
	}

	return repos
}

func (s Service) Create(name string, description string) (Repo, error) {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return Repo{}, err
	}

	repoPath := path.Join(pwd, "repos", fmt.Sprintf("%s.git", name))

	cmd := exec.Command("mkdir", "-p", repoPath)
	err = cmd.Run()
	if err != nil {
		log.Printf("Failed to create directory: %v", repoPath)
		return Repo{}, err
	}

	cmd = exec.Command("git", "init", "--bare", "--initial-branch", "main")
	cmd.Dir = repoPath
	err = cmd.Run()
	if err != nil {
		log.Printf("Failed to initialize git repository: %v", err)
		return Repo{}, err
	}

	cmd = exec.Command("touch", "git-daemon-export-ok")
	cmd.Dir = repoPath
	err = cmd.Run()
	if err != nil {
		log.Printf("Failed to create git-daemon-export-ok: %v", err)
		return Repo{}, err
	}

	repo := Repo{
		Name:        name,
		Description: strings.TrimSpace(description),
		Url:         fmt.Sprintf("%s/%s.git", s.Url, name),
	}

	return repo, nil
}

func (s Service) Update() {

}

func (s Service) Destroy() {

}

func (s Service) Get() {

}

func TrimSuffix(s, suffix string) string {
	if strings.HasSuffix(s, suffix) {
		s = s[:len(s)-len(suffix)]
	}
	return s
}
