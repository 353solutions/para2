package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var (
	//go:embed html/last.html
	lastHTML string
)

type Server struct {
	db *DB
}

func (s *Server) lastHandler(w http.ResponseWriter, r *http.Request) {
	lastText := "No entries"

	entry, err := s.db.Last()
	if err == nil {
		time := entry.Time.Format("2006-01-02T15:04")
		lastText = fmt.Sprintf("[%s] %s by %s", time, entry.Content, entry.User)
	}
	fmt.Fprintf(w, lastHTML, lastText)
}

func (s *Server) newHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var e Entry

	if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	e.Time = time.Now()
	if err := s.db.Add(e); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(e)
}

func (s *Server) Health() error {
	return s.db.Health()
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	if err := s.Health(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Fprintf(w, "OK\n")
}

func main() {
	var err error
	db, err := NewDB("host=localhost user=postgres password=s3cr3t sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	s := Server{db}

	r := mux.NewRouter()
	r.HandleFunc("/last", s.lastHandler).Methods(http.MethodGet)
	r.HandleFunc("/new", s.newHandler).Methods(http.MethodPost)
	r.HandleFunc("/health", s.healthHandler).Methods(http.MethodGet)

	http.Handle("/", r)

	const addr = ":8080"
	log.Printf("server starting on %s", addr)
	err = http.ListenAndServe(":8080", nil)
	s.db.Close()
	if err != nil {
		log.Fatal(err)
	}
}
