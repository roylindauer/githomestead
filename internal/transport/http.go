package transport

import (
	"encoding/json"
	"gitapi/internal/gitserver"
	"log"
	"net/http"
)

type Server struct {
	mux *http.ServeMux
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

		w.WriteHeader(http.StatusOK)
	})

	mux.HandleFunc("POST /repos", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		repo, err := gitSvc.Create("itsarepo", "it code")
		if err != nil {
			log.Fatal(err)
		}

		b, err := json.Marshal(repo)
		if err != nil {
			log.Fatal(err)
		}

		_, err = w.Write(b)
		if err != nil {
			log.Fatal(err)
		}

		w.WriteHeader(http.StatusOK)
	})

	return &Server{
		mux: mux,
	}
}

func (s *Server) Serve() error {
	return http.ListenAndServe(":8080", s.mux)
}
