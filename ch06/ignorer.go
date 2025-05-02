package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	fmt.Println("program started")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for sig := range sigChan {
			fmt.Printf("received signal: %v (ignoring)\n", sig)
		}
	}()

	for {
		fmt.Println("program is running...")
		time.Sleep(1 * time.Second)
	}
}
