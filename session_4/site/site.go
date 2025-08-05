package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	users := http.NewServeMux()
	// We don't want users to know where's it's mounted on top router
	// users.HandleFunc("GET /users/{uid}", handler)
	users.HandleFunc("GET /{uid}", handler)

	admin := http.NewServeMux()
	admin.HandleFunc("POST /shutdown", handler)

	root := http.NewServeMux()
	root.Handle("/users/", http.StripPrefix("/users", users))
	root.Handle("/admin/", http.StripPrefix("/admin", admin))

	srv := http.Server{
		Handler: root,
		Addr:    ":9999",
	}
	if err := srv.ListenAndServe(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}

}

/*
Code without http.HandleFunc and http.HandlerFunc
func (adminShutdownHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {}

type adminShutdownHandler struct{}

func (getUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {}

type getUserHandler struct{}
*/

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "handler: %s\n", r.URL.Path)
}
