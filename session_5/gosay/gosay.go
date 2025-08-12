package main

import (
	"embed"
	"flag"
	"fmt"
	"os"
	"path"
	"strings"
)

var version = "<dev>"

var options struct {
	showVersion bool
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

func main() {
	flag.BoolVar(&options.showVersion, "version", false, "show version & exit")
	flag.Parse()

	if options.showVersion {
		fmt.Printf("%s version %s\n", path.Base(os.Args[0]), version)
	}

	if flag.NArg() != 1 {
		fmt.Fprintln(os.Stderr, "error: wrong number of arguments")
		os.Exit(1)
	}

	text := flag.Arg(0)
	width := len(text)
	fmt.Printf(" %s\n", strings.Repeat("-", width))
	fmt.Printf("< %s >\n", text)
	fmt.Printf(" %s\n", strings.Repeat("-", width))
	fmt.Println(gopher)
}

// $ go build -ldflags='-X main.version=0.1.0'
// $ GOOS=windows GOARCH=amd64 go build -ldflags="-H windowsgui"
