package main

import (
	"compress/gzip"
	"crypto/sha1"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func main() {
	fmt.Println(FileSHA1("http.log.gz"))
	fmt.Println(FileSHA1("road.txt"))
}

// Exercise: Only decompress if file name ends with .gz
// cat http.log.gz| gunzip |sha1sum
// cat road.txt| sha1sum
func FileSHA1(fileName string) (string, error) {
	// cat http.log.gz
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var r io.Reader = file

	if path.Ext(fileName) == ".gz" {
		// | gunzip
		// r, err := gzip.NewReader(file) // BUG: shadows "r" on line 28
		gz, err := gzip.NewReader(file)
		if err != nil {
			return "", fmt.Errorf("gzip %q - %w", fileName, err)
		}
		defer gz.Close()
		r = gz
	}

	// |sha1sum
	w := sha1.New()
	if _, err := io.Copy(w, r); err != nil {
		return "", fmt.Errorf("copy %q - %w", fileName, err)
	}

	sig := w.Sum(nil)
	return fmt.Sprintf("%x", sig), nil
}

type gzWriter struct {
	http.ResponseWriter
	out *gzip.Writer
}

func (gz *gzWriter) Write(data []byte) (int, error) {
	return gz.out.Write(data)
}

func gzMiddleWare(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Accept-Encoding") == "gzip" {
			w = &gzWriter{w, gzip.NewWriter(w)}
			w.Header().Set("Content-Encoding", "gzip")
		}

		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func gzHandler(w http.ResponseWriter, r *http.Request) {
	var out io.Writer = w
	if r.Header.Get("Accept-Encoding") == "gzip" {
		w.Header().Set("Content-Encoding", "gzip")
		out = gzip.NewWriter(w)
	}

	fmt.Fprintln(out, "Hello")
}

/* Thought experiment: Sorting

type Sortable interface {
	Less(i, j int) bool
	Swap(i, j)
	Len() int
}

func Sort(s Sortable) {
	// sort s
}

*/
