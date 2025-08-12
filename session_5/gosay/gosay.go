package main

import (
	"embed"
	"flag"
	"fmt"
	"io"
	"os"
	"path"
	"strings"
)

var version = "<dev>"

var options struct {
	showVersion bool
	image       string
}

// Most import embed for this to work
// Can embed into string or []byte
//
//go:embed gopher.txt
var gopher string

//go:embed images
var images embed.FS

// Exercise: Support -image NAME that will use image from the "images"
// directory instead of the gopher
// gosay -image llama

// "go build" does not run "go generate"
//go:generate bash -c "echo hello > greeting.txt"

func main() {
	flag.BoolVar(&options.showVersion, "version", false, "show version & exit")
	flag.StringVar(&options.image, "image", "", "image to use")
	flag.Parse()

	if options.showVersion {
		fmt.Printf("%s version %s\n", path.Base(os.Args[0]), version)
	}

	var text string
	switch flag.NArg() {
	case 0:
		data, err := io.ReadAll(io.LimitReader(os.Stdin, 1024))
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: read stdin - %s\n", err)
			os.Exit(1)
		}
		text = string(data)
	case 1:
		text = flag.Arg(0)
	default:
		fmt.Fprintln(os.Stderr, "error: wrong number of arguments")
		os.Exit(1)
	}

	width := len(text)
	fmt.Printf(" %s\n", strings.Repeat("-", width))
	fmt.Printf("< %s >\n", text)
	fmt.Printf(" %s\n", strings.Repeat("-", width))
	if options.image == "" {
		fmt.Println(gopher)
		return
	}

	file, err := images.Open(fmt.Sprintf("images/%s.txt", options.image))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s - %s\n", options.image, err)
		os.Exit(1)
	}
	defer file.Close()
	io.Copy(os.Stdout, file)
}

// $ go build -ldflags='-X main.version=0.1.0'
// $ GOOS=windows GOARCH=amd64 go build -ldflags="-H windowsgui"
