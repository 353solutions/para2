package main

import (
	"context"
	"embed"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
	"time"
)

var version = "<dev>"

var options struct {
	showVersion bool
	image       string
	useCow      bool
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

/*
	Nesting in environment

	GOSAY_LOG__LEVEL: INFO
	GOSAY_LOG__FILE_NAME: /var/log/gosay.log
	GOSAY_SERVER__PORT: 9090
	GOSAY_SERVER__VERBOSE: false
	GOSAY_HOME: /local/gosay


	{
		Log: {
			Level: "INFO",
		},
		Server: {
			Port: 8080,
			Verbose: false,
		}
	}
*/

func main() {
	flag.BoolVar(&options.showVersion, "version", false, "show version & exit")
	flag.StringVar(&options.image, "image", "", "image to use")
	flag.BoolVar(&options.useCow, "cow", false, "use `cow` instead")
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

	// All options should be parsed and validated here
	// Example: -addr :8080
	// - check for ':' in address
	// - check that part after : is an int between 0 to 65_535

	if options.useCow {
		if err := cow(text); err != nil {
			fmt.Fprintf(os.Stderr, "error: using cow - %s\n", err)
			os.Exit(1)
		}
		return
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
// $ CGO_ENABLED=0 go build

/* Configuration
defaults < configuration file < environment < command line
defaults:
	code
configuration file:
	not in stdlib
	3rdParty: YAML, TOML, viper
environment:
	os.Getenv (stdlib)
	viper
command line:
	flag (stdlib)
	cobra

See also ardanlabs/conf

Most important: Validate configuration *before* running
*/

func cow(text string) error {
	q := url.Values{}
	q.Add("message", text)
	q.Add("format", "text")
	url := "https://cowsay.morecode.org/say?" + q.Encode()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%q - %s", url, resp.Status)
	}

	_, err = io.Copy(os.Stdout, resp.Body)
	fmt.Println()
	return err
}

/*
$ ldd gosway
        linux-vdso.so.1 (0x00007fd970f0e000)
        libresolv.so.2 => /usr/lib/libresolv.so.2 (0x00007fd970ec4000)
        libc.so.6 => /usr/lib/libc.so.6 (0x00007fd970c00000)
        /lib64/ld-linux-x86-64.so.2 => /usr/lib64/ld-linux-x86-64.so.2 (0x00007fd970f10000)
*/
