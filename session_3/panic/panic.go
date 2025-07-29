package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

/* Use named return values for:
- panic
- documentation
*/

// recover works only for the current goroutine

func main() {
	fmt.Println(safeDiv(7, 2))
	fmt.Println(safeDiv(1, 0))

	httpd()
}

func handler(w http.ResponseWriter, r *http.Request) {
	slog.Info("handler")

	/* Will crash the server
	go func() {
		v := div(1, 0) // panic
		slog.Info("add result", "v", v)
	}()
	*/
	safelyGo(func() {
		v := div(1, 0) // panic
		slog.Info("add result", "v", v)
	})

	fmt.Fprintf(w, "OK")
}

func safelyGo(fn func()) {
	go func() {
		defer func() {
			if e := recover(); e != nil {
				slog.Warn("goroutine panic", "error", e)
			}
		}()

		fn()

	}()

}

func httpd() {
	http.HandleFunc("/", handler)

	addr := ":8081"
	slog.Info("server starting", "address", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}

// Don't use this style
func add(a, b int) (out int) {
	out = a + b
	// ....
	return
}

// named return value
func safeDiv(a, b int) (q int, err error) {
	// q & err are variables in safeDiv (like a & b)
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e) // e is any, need to convert to error
		}
	}()

	return div(a, b), nil
}

func div(a, b int) int {
	return a / b
}
