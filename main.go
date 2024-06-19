package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

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
	signal := <-sigChan

	fmt.Println(signal)

	// cancel ctx
	cancel()

	time.Sleep(1 * time.Second)

}
