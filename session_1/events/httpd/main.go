package main

import (
	"encoding/json"
	"flag"
	"io"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"time"
)

var config struct {
	auth bool
}

func main() {
	flag.BoolVar(&config.auth, "auth", false, "use authentication")
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("POST /events", handleEvents)

	const addr = ":8080"
	srv := http.Server{
		Handler: mux,
		Addr:    addr,
	}

	slog.Info("server starting", "address", addr)
	if err := srv.ListenAndServe(); err != nil {
		slog.Error("listen", "error", err)
		os.Exit(1)
	}
}

func handleEvents(w http.ResponseWriter, r *http.Request) {
	slog.Info("events")

	if config.auth {
		hdr := r.Header.Get("Authorization")
		hdr = strings.TrimPrefix(hdr, "Bearer ")
		if hdr != "s3cr3t" {
			slog.Error("login", "remote", r.RemoteAddr)
			http.Error(w, "bad login", http.StatusUnauthorized)
			return
		}
	}

	defer r.Body.Close()
	rdr := io.TeeReader(r.Body, os.Stdout)
	dec := json.NewDecoder(rdr)
	count := 0
	for {
		var e Event
		err := dec.Decode(&e)
		if err == io.EOF {
			break
		}
		if err != nil {
			slog.Error("bad json", "error", err)
			http.Error(w, "bad JSON", http.StatusBadRequest)
			return
		}
		count++
	}

	slog.Info("events", "count", count)

	resp := map[string]any{
		"count": count,
	}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		slog.Warn("response", "error", err)
	}
}

type Event struct {
	Time   time.Time
	Login  string
	Action string
	URI    string
}
