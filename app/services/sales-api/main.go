package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/rahjooh/service/foundation/logger"
	"go.uber.org/zap"
)

var build = "develop"

func main() {
	// Construct the application logger.
	log, err := logger.New("SALES-API")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer log.Sync()

	// Perform the startup and shutdown sequence.
	if err := run(log); err != nil {
		log.Errorw("startup", "ERROR", err)
		os.Exit(1)
	}
}
func run(log *zap.SugaredLogger) error {

	// =========================================================================
	// GOMAXPROCS

	log.Infow("startup", "GOMAXPROCS", runtime.GOMAXPROCS(0), "BUILD", build)

	// =========================================================================
	// GOMAXPROCS

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	sig := <-shutdown

	log.Infow("shutdown", "status", "shutdown started", "signal", sig)
	defer log.Infow("shutdown", "status", "shutdown compelete", "signal", sig)

	return nil
}
