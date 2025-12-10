package app

import (
	"context"
	"fmt"
	"os"

	"github.com/mrxacker/go-myapp/internal/config"
	"github.com/mrxacker/go-myapp/pkg/logger/zap"
)

func Run(ctx context.Context) error {
	_, err := config.Load()
	if err != nil {
		return err
	}

	logger, cleanup, err := zap.NewLogger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing logger: %v\n", err)
		os.Exit(1)
	}

	defer cleanup()

	logger.Info("Starting application")

	return nil
}

func runServers(ctx context.Context) error {

	// Placeholder for server initialization and running logic
	return nil
}
