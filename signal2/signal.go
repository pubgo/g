package signal2

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

var onlyOneSignalHandler = make(chan struct{})

// SetupSignalHandler
// registered for SIGTERM and SIGINT. A stop channel is returned
// which is closed on one of these signals. If a second signal is caught, the program
// is terminated with exit code 1.
func SetupSignalHandler() (stopCh <-chan struct{}) {
	close(onlyOneSignalHandler) // panics when called twice

	stop := make(chan struct{})
	c := make(chan os.Signal, 2)
	signal.Notify(c, shutdownSignals...)
	go func() {
		<-c
		close(stop)
		<-c
		os.Exit(1) // second signal. Exit directly.
	}()

	return stop
}

// ErrSignalTrapped is returned by the SignalTrap.Wait
// when the expected signals caught.
const ErrSignalTrapped = "signal trapped"

// Termination returns trap for termination signals.
//
//  server := new(http.Server)
//  go log.Println(server.ListenAndServe())
//
//  err := sync.Termination().Wait(context.Background())
//  if err == sync.ErrSignalTrapped {
//  	log.Println("shutting down the server", server.Shutdown(context.Background()))
//  }
//
func Termination() SignalTrap {
	trap := make(chan os.Signal, 3)
	signal.Notify(trap, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	return trap
}

// SignalTrap wraps os.Signal channel to provide high level API above it.
//
//  trap := make(chan os.Signal)
//  signal.Notify(trap, os.Interrupt)
//  SignalTrap(trap).Wait(context.Background())
//
type SignalTrap chan os.Signal

// Wait blocks until one of the expected signals caught
// or the Context closed. It unregisters from the notification
// and closes itself.
func (trap SignalTrap) Wait(ctx context.Context) error {
	defer close(trap)
	defer signal.Stop(trap)

	select {
	case <-trap:
		return fmt.Errorf(ErrSignalTrapped)
	case <-ctx.Done():
		return ctx.Err()
	}
}
