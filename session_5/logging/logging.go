package main

import (
	"fmt"
	"os"

	"go.uber.org/zap"
)

/*
Log in human format to stdout and JSON to log file
In docker: Write JSON to stdout

In general set logger from configuration, do it first.
*/

func main() {
	log, err := zap.NewDevelopment()
	// log, err := zap.NewProduction()
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

}

/*
If logging doesn't help you during development, it won't help you at 4am when production is burning.
	- Bill Kennedy
*/
