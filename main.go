package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/dbtedman/scrutinise/infrastructure/cmd"
)

func main() {
	signalsCh := make(chan os.Signal, 1)
	resultCh := make(chan int)
	var result int

	// listen to os signals
	signal.Notify(signalsCh, os.Interrupt, syscall.SIGTERM)

	defer func() {
		// cleanup any resources now, then exit
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
