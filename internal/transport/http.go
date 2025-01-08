package transport

import (
	"encoding/json"
	"fmt"
	"gitapi/internal/gitserver"
	"log"
	"net/http"
	"strings"
)

type Server struct {
	mux *http.ServeMux
}

type Repo struct {
	Name        string
	Description string
}

func NewServer(gitSvc gitserver.Service) *Server {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /repos", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		b, err := json.Marshal(gitSvc.GetAllRepos())
		if err != nil {
			log.Fatal(err)
		}

		_, err = w.Write(b)
		if err != nil {
			log.Fatal(err)
		}
	})

	mux.HandleFunc("POST /repos", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var repoItem Repo

		err := json.NewDecoder(r.Body).Decode(&repoItem)

		err = gitSvc.Create(repoItem.Name, repoItem.Description)
		if err != nil {
			log.Fatal(err)
		}

		repo := gitSvc.Get(fmt.Sprintf("%s.git", repoItem.Name))

		b, err := json.Marshal(repo)
		if err != nil {
			log.Fatal(err)
		}

		_, err = w.Write(b)
		if err != nil {
			log.Fatal(err)
		}
	})

	mux.HandleFunc("GET /repo/{name}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		b, err := json.Marshal(gitSvc.Get(r.PathValue("name")))
		if err != nil {
			log.Fatal(err)
		}

		_, err = w.Write(b)
		if err != nil {
			log.Fatal(err)
		}
	})

	return &Server{
		mux: mux,
	}
}

func (s *Server) Serve() error {
	return http.ListenAndServe(":8080", s.mux)
}

func TrimSuffix(s, suffix string) string {
	if strings.HasSuffix(s, suffix) {
		s = s[:len(s)-len(suffix)]
	}
	return s
}
