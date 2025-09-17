package main

import (
	"fmt"
	"io"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

/*
Log in human format to stdout and JSON to log file
In docker: Write JSON to stdout

In general set logger from configuration, do it first.
*/

func main() {
	// log, err := zap.NewDevelopment()
	// log, err := zap.NewProduction()
	log, err := createLogger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
	defer log.Sync()

	slog := log.Sugar()
	// Argument go via heap since port and int are "any"
	slog.Infow("system up", "port", 9999)

	// No escape to the heap
	log.Info("connect to database",
		zap.String("db_host", "localhost"),
		zap.Int("db_port", 7654),
	)

	log.Debug("in main")

	start := time.Now()
	if log.Core().Enabled(zapcore.DebugLevel) {
		log.Debug("user", zap.String("login", CurrentUser()))
	}
	duration := time.Since(start)
	fmt.Println(duration)

	log.Info("user role", zap.Any("role", Writer))

	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
	defer file.Close()

	w := io.MultiWriter(os.Stdout, file)
	fmt.Fprintln(w, "much writing")
}

//go:generate go tool stringer -type Role
// Sometimes fmt.Stringer is not enough and you'll want also to
// implement text.Marshaler or json.Marshaler

const (
	Reader Role = iota + 1
	Writer
	Admin
)

type Role byte

func createLogger() (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	// cfg := zap.NewDevelopmentConfig()
	cfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	cfg.OutputPaths = []string{"app.log", "stdout"}
	cfg.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder

	log, err := cfg.Build()
	if err != nil {
		return nil, err
	}

	log = log.With(zap.String("host", os.Getenv("HOST")))
	return log, nil
}

func CurrentUser() string {
	time.Sleep(time.Second)
	return "scofield"
}

/*
If logging doesn't help you during development, it won't help you at 4am when production is burning.
	- Bill Kennedy

$ go build -trimpath
	- Will make files relative to module root
*/
