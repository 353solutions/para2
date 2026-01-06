package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	log := slog.Default().With("app", "unter")
	dbFile := os.Getenv("UNTER_DB")
	if dbFile == "" {
		dbFile = "unter.db"
	}
	addr := os.Getenv("UNTER_ADDR")
	if addr == "" {
		addr = ":8080"
	}

	db, err := NewDB(dbFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: can't create DB (%s)\n", err)
		os.Exit(1)
	}

	api := API{
		log: log,
		db:  db,
	}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", api.Health)
	mux.HandleFunc("POST /rides", api.Add)
	mux.HandleFunc("GET /ride/{id}", api.Get)

	srv := http.Server{
		Addr:    addr,
		Handler: mux,
	}
	log.Info("server starting", "address", srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Error("serve", "error", err)
		os.Exit(1)
	}
}
