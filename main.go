package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/dbtedman/scrutinise/infrastructure/cmd"
)

func main() {
	logger := cmd.Logger(os.Stdout)
	logger.Info("scrutinise :: begin")

	signalsCh := make(chan os.Signal, 1)
	resultCh := make(chan int)
	var result int

	// listen to os signals
	signal.Notify(signalsCh, os.Interrupt, syscall.SIGTERM)

	defer func() {
		// cleanup any resources now, then exit
		logger.Info("scrutinise :: end")
		os.Exit(result)
	}()

	go func() {
		err := cmd.RootCommand(&resultCh).Execute()

		if err != nil {
			fmt.Println(err)
			resultCh <- cmd.ErrorResult
		}
	}()

	select {
	case <-signalsCh:
	case result = <-resultCh:
	}
}
