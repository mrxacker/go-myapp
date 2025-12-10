package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/mrxacker/go-myapp/internal/app"
)

func main() {
	log.Println("Starting MyApp Service")

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	if err := app.Run(ctx); err != nil {
		log.Fatalf("app failed to run: %v", err)
	}
}
