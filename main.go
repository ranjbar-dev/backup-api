package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/ranjbar-dev/backup-api/srv/api"
)

func main() {

	// Create a context with cancellation
	ctx, cancel := context.WithCancel(context.Background())

	// Create a channel to receive OS signals
	sigChan := make(chan os.Signal, 1)

	// Notify the sigChan on receiving signals
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// start api
	a := api.NewApi(ctx)
	a.Start()

	// wait for signal
	<-sigChan

	// cancel ctx
	cancel()

}
