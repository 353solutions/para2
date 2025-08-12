package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	_ "net/http/pprof"
	"os"
	"solutions/collatz"
	"strconv"
)

func main() {
	http.HandleFunc("GET /max/{n}", maxHandler)

	addr := ":8081"
	slog.Info("server starting", "address", addr, "pid", os.Getpid())
	if err := http.ListenAndServe(addr, nil); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}

func maxHandler(w http.ResponseWriter, r *http.Request) {
	n, err := strconv.Atoi(r.PathValue("n"))
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	if n <= 0 {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	m, l := collatz.MaxCollatz(n)
	resp := map[string]int{
		"n":   m,
		"len": l,
	}
	w.Header().Set("content-type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		slog.Error("can't encode", "error", err)
	}
}
