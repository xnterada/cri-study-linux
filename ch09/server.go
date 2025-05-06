package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: ./server <port_number>")
		os.Exit(1)
	}

	portStr := os.Args[1]
	port, err := strconv.Atoi(portStr)
	if err != nil {
		fmt.Printf("invalid port number: %s\n", portStr)
		os.Exit(1)
	}

	http.Handle("/", http.FileServer(http.Dir("./html")))

	addr := fmt.Sprintf(":%d", port)
	log.Printf("Starting server on port %s...", addr)
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}