package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
)

type API struct {
	log *slog.Logger
	db  *DB
}

func (a *API) Health(w http.ResponseWriter, r *http.Request) {
	// TODO: Health check
	fmt.Fprintln(w, "OK")
}

func (a *API) Get(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "missing ID", http.StatusBadRequest)
		return
	}

	rd, err := a.db.Get(id)
	if err != nil {
		a.log.Error("scan", "error", err)
		http.Error(w, "can't get rides", http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	if err := json.NewEncoder(w).Encode(rd); err != nil {
		a.log.Error("encode", "error", err)
		return
	}
}

func (a *API) Add(w http.ResponseWriter, r *http.Request) {
	var rd Ride
	if err := json.NewDecoder(r.Body).Decode(&rd); err != nil {
		a.log.Error("decode", "error", err)
		http.Error(w, "bad record", http.StatusBadRequest)
		return
	}

	if err := rd.Validate(); err != nil {
		a.log.Error("validate", "error", err)
		http.Error(w, "bad record", http.StatusBadRequest)
		return
	}

	if err := a.db.Insert(rd); err != nil {
		a.log.Error("insert", "error", err)
		http.Error(w, "can't insert", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]any{
		"id": rd.ID,
	})
}
