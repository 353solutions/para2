package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	/*
		w := bufio.NewWriter(os.Stdout)
		log := NewLogger(w)
	*/
	log := NewLogger(os.Stdout)
	log.Info("server starting on port %d", 8080)
}

func (l *Logger) Info(format string, args ...any) {
	// args is []any
	fmt.Fprint(l.w, "INFO: ")
	fmt.Fprintf(l.w, format, args...)
	l.f.Flush()
}

func NewLogger(w io.Writer) *Logger {
	l := Logger{
		w: w,
		f: NullFlusher{},
	}
	// Check at runtime that w underlying type implements interface
	if f, ok := l.w.(flusher); ok {
		l.f = f
	} else if s, ok := l.w.(syncer); ok {
		l.f = &SyncFlusher{s}
	}

	return &l
}

func (s *SyncFlusher) Flush() error {
	return s.s.Sync()
}

type SyncFlusher struct {
	s syncer
}
type syncer interface {
	Sync() error // *os.File
}

func (NullFlusher) Flush() error { return nil }

type NullFlusher struct{}

type flusher interface {
	Flush() error
}

type Logger struct {
	w io.Writer
	f flusher
	// Not using it since limits the amount of types the logger can accept as w
	// w WriteFlusher
}

type WriteFlusher interface {
	io.Writer
	Flush() error
}
