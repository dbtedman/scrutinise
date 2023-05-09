package main

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

const errorResult = 1

func main() {
	fmt.Println("scrutinise")

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
		err := errors.New("no command is defined yet")

		if err != nil {
			fmt.Println(err)
			resultCh <- errorResult
		}
	}()

	select {
	case <-signalsCh:
	case result = <-resultCh:
	}
}
