package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"os/signal"
	"path"
	"strings"
	"syscall"
)

func main() {
	showLine := false
	flag.BoolVar(&showLine, "n", false, "show line numbers")
	// flag.BoolVar(&showLine, "number", false, "show line numbers")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: %s TERM FILE\n", path.Base(os.Args[0]))
		flag.PrintDefaults()
	}
	flag.Parse()

	var r io.Reader
	switch flag.NArg() {
	case 1:
		r = os.Stdin
	case 2:
		fileName := flag.Arg(1)
		file, err := os.Open(fileName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n", err)
			os.Exit(1)
		}
		defer file.Close()
		r = file
	default:
		fmt.Fprintln(os.Stderr, "error: wrong number of arguments")
		os.Exit(1)
	}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGPIPE)
	go func() {
		<-sigCh
		fmt.Fprintln(os.Stderr, "PIPE")
		os.Exit(0)
	}()
	// defer fmt.Fprintln(os.Stderr, "PIPE")

	term := flag.Arg(0)
	s := bufio.NewScanner(r)
	lNum := 0
	for s.Scan() {
		lNum++
		line := s.Text()
		if !strings.Contains(line, term) {
			continue
		}

		/* TODO: Doesn't show well n logs ...
		Solutions:
		1. Ask user for a flag
		2. Detect if running in terminal (isatty)
		*/
		line = mark(term, line)

		if showLine {
			// TODO: How do we know padding
			fmt.Printf("%3d: %s\n", lNum, line)
		} else {
			fmt.Println(line)
		}
	}
	if err := s.Err(); err != nil {
		// TODO Add context
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)

	}
}

func mark(expr, line string) string {
	// TODO: Support more than one expr in line
	i := strings.Index(line, expr)
	if i == -1 {
		return line
	}
	j := i + len(expr)

	return line[:i] + "\033[31m" + expr + "\033[39m" + line[j:]
}

/* In shell
CTRL-D: EOF
CTRL-C: Interrupt
*/

/* Assume not all input will be read
You might get SIGPIPE signal

3 main signals:
SIGPIPE: stdout/stderr closed
SIGINT: CTRL-C
SIGHUP: Terminal closed
	- see nohup

You can't catch all signals (e.g. KILL=9)
*/
