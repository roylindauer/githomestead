package gitserver

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

func Endpoint() string {
	return "git://localhost:9418"
}

type Service struct {
	Url      string
	RepoPath string
}

func NewService() *Service {
	return &Service{
		Url:      Endpoint(),
		RepoPath: RepoDir(),
	}
}

type Repo struct {
	Url           string
	Description   string
	Name          string
	GitName       string
	DefaultBranch string
	repoPath      string
}

func NewRepo(name string) Repo {
	normalizedName := TrimSuffix(name, ".git")
	normalizedName = Slugify(normalizedName)
	repo := Repo{
		Name:          normalizedName,
		GitName:       fmt.Sprintf("%s.git", normalizedName),
		DefaultBranch: "main",
		Description:   "",
	}
	repo.Url = fmt.Sprintf("%s/%s", Endpoint(), repo.GitName)
	repo.repoPath = path.Join(RepoDir(), repo.GitName)

	return repo
}

func (s Service) GetAllRepos() []Repo {
	var repos []Repo

	directories, err := os.ReadDir(s.RepoPath)

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range directories {
		if file.IsDir() && strings.HasSuffix(file.Name(), ".git") {
			repos = append(repos, s.Get(file.Name()))
		}
	}

	return repos
}

func RootDir() string {
	pwd, err := os.Getwd()
	if err != nil {
		log.Printf("Failed to get root directory: %v", err)
		panic(err)
	}

	return pwd
}

func RepoDir() string {
	return path.Join(RootDir(), "repos")
}

func (s Service) RepoExists(repoPath string) bool {
	_, err := os.Open(repoPath)
	if os.IsExist(err) {
		return true
	}
	return false
}

func (s Service) Create(name string, description string) error {
	repo := NewRepo(name)
	repo.Description = description

	if s.RepoExists(repo.repoPath) {
		return nil
	}

	cmd := exec.Command("mkdir", "-p", repo.repoPath)
	err := cmd.Run()
	if err != nil {
		log.Printf("Command: %v", cmd)
		log.Printf("Failed to create directory: %v", repo.repoPath)
		return err
	}

	cmd = exec.Command("git", "init", "--bare", "--initial-branch", repo.DefaultBranch)
	cmd.Dir = repo.repoPath
	err = cmd.Run()
	if err != nil {
		log.Printf("Command: %v", cmd)
		log.Printf("Failed to initialize git repository: %v", err)
		return err
	}

	cmd = exec.Command("touch", "git-daemon-export-ok")
	cmd.Dir = repo.repoPath
	err = cmd.Run()
	if err != nil {
		log.Printf("Failed to create git-daemon-export-ok: %v", err)
		return err
	}

	if description != "" {
		if err := os.WriteFile(path.Join(repo.repoPath, "description"), []byte(description), 0644); err != nil {
			log.Fatal(err)
		}
	}

	return nil
}

func (s Service) Update() {

}

func (s Service) Destroy() {

}

func (s Service) Get(repoName string) Repo {
	repo := NewRepo(repoName)

	description, err := os.ReadFile(filepath.Join(repo.repoPath, "description"))
	repo.Description = strings.TrimSpace(string(description))

	if err != nil {
		log.Fatal(err)
	}

	return repo
}

// GetCommits What if we Return array of commits for the repo
func (s Service) GetCommits(repo Repo) {

}

func TrimSuffix(s, suffix string) string {
	if strings.HasSuffix(s, suffix) {
		s = s[:len(s)-len(suffix)]
	}
	return s
}

func Slugify(s string) string {
	result := strings.ToLower(s)
	result = strings.ReplaceAll(result, " ", "-")
	reg, err := regexp.Compile("[^a-z0-9-]+")
	if err != nil {
		fmt.Println(err)
		return ""
	}
	result = reg.ReplaceAllString(result, "")
	return result
}
