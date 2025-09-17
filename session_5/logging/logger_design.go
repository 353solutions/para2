package main

import (
	"io"
	"time"
)

type Level int

const (
	InfoLevel Level = iota + 1
)

type Record struct {
	Time    time.Time
	Level   Level
	Message string
	Attrs   map[string]any
}

type Formatter struct {
}

func (f Formatter) Format(r Record) string {
	return ""
}

type Handler struct {
	w io.Writer
	f Formatter
}

func (h Handler) Handle(r Record) {
	s := h.f.Format(r)
	h.w.Write([]byte(s))
}

type Logger struct {
	Level Level

	handlers []Handler
}

func (l Logger) Info(msg string) {
	if l.Level < InfoLevel {
		return
	}

	r := Record{
		Time:    time.Now(),
		Message: msg,
		// ...
	}

	for _, h := range l.handlers {
		h.Handle(r)
	}
}
